// resource_server.go
package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func organisationData() *schema.Resource {
	return &schema.Resource{
		Read: resourceServerRead,
	}
}

func StateFunc(data *schema.ResourceData, iface interface{}) ([]*schema.ResourceData, error) {
	retVal := make([]*schema.ResourceData, 1)
	err := data.Set(OrganisationId, data.Id())

	if err != nil {
		return nil, err
	}

	err = resourceServerRead(data, iface)
	if err != nil {
		return nil, err
	}

	retVal[0] = data
	return retVal, nil
}

func organisationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{State: StateFunc},
		Create:   organisationCreate,
		Read:     resourceServerRead,
		Update:   resourceServerUpdate,
		Delete:   resourceServerDelete,

		Schema: map[string]*schema.Schema{
			OrganisationId: {
				Type:     schema.TypeString,
				Required: true,
			},
			"organisation_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func organisationCreate(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	organisationName := d.Get("organisation_name").(string)

	response, err := m.(SafeToRunProvider).Client.CreateOrganisationAndWait(safetorun.CreateOrganisationRequest{
		OrganisationId:   organisationId,
		OrganisationName: organisationName,
	})

	if err != nil {
		return err
	}

	d.SetId(response.OrganisationId)

	if err != nil {
		return err
	}

	err = d.Set("organisation_name", organisationName)

	if err != nil {
		return err
	}

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	d.SetId(organisationId)
	_, err := m.(SafeToRunProvider).Client.QueryStatus(organisationId)
	return err
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	response, err := m.(SafeToRunProvider).Client.DeleteOrganisationAndWait(organisationId)
	if err != nil {
		log.Fatal("failed to delete", err)
		return err
	}
	d.SetId(response.OrganisationId)

	return nil
}

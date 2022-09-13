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

func organisationResource() *schema.Resource {
	return &schema.Resource{
		Create: organisationCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func organisationCreate(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get("organisation_id").(string)
	println("Going to create")
	response, err := m.(SafeToRunProvider).Client.CreateOrganisation(safetorun.CreateOrganisationRequest{OrganisationId: organisationId})

	println("Created")
	if err != nil {
		log.Fatal("failed to create", err)
		return err
	}

	d.SetId(response.OrganisationId)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

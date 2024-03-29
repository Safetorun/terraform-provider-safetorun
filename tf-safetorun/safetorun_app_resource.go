package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func StateFuncApp(data *schema.ResourceData, iface interface{}) ([]*schema.ResourceData, error) {
	retVal := make([]*schema.ResourceData, 1)

	id := strings.Split(data.Id(), ".")
	err := data.Set(OrganisationId, id[0])
	if err != nil {
		return nil, err
	}

	err = data.Set(ApplicationId, id[1])

	if err != nil {
		return nil, err
	}

	err = applicationRead(data, iface)
	if err != nil {
		return nil, err
	}

	retVal[0] = data
	return retVal, nil
}

func applicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: StateFuncApp,
		},
		Create: applicationCreate,
		Read:   applicationRead,
		Update: applicationUpdate,
		Delete: applicationDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			OrganisationId: {
				Type:     schema.TypeString,
				Required: true,
			},
			ApplicationId: {
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"api_key": {
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
		},
	}
}

func applicationUpdate(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	applicationId := d.Get(ApplicationId).(string)
	applicationName := d.Get("application_name").(string)

	_, err := m.(SafeToRunProvider).Client.UpdateApplication(safetorun.UpdateApplicationRequest{
		OrganisationId:  organisationId,
		ApplicationId:   applicationId,
		ApplicationName: applicationName,
	})

	if err != nil {
		return err
	}

	return applicationRead(d, m)
}

func applicationRead(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	applicationId := d.Get(ApplicationId).(string)
	d.SetId(applicationId)
	response, err := m.(SafeToRunProvider).Client.QueryApplication(organisationId, applicationId)

	if err != nil {
		return err
	}

	err = d.Set(ApplicationId, response.ApplicationId)

	if err != nil {
		return err
	}

	err = d.Set("api_key", response.ApiKey)
	if err != nil {
		return err
	}

	err = d.Set("application_name", response.ApplicationName)
	return err
}

func applicationDelete(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	applicationName := d.Get("application_name").(string)

	_, err := m.(SafeToRunProvider).Client.DeleteApplication(safetorun.DeleteApplicationRequest{
		OrganisationId: organisationId,
		ApplicationId:  applicationName,
	})

	if err != nil {
		return nil
	}

	return nil
}

func applicationCreate(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	applicationName := d.Get("application_name").(string)

	response, err := m.(SafeToRunProvider).Client.CreateApplicationAndWait(safetorun.CreateApplicationRequest{
		OrganisationId:  organisationId,
		ApplicationName: applicationName,
	})

	if err != nil {
		log.Fatal("failed to create", err)
		return err
	}

	d.SetId(response.ApplicationId)
	err = d.Set("application_name", applicationName)

	if err != nil {
		return err
	}

	err = d.Set(ApplicationId, response.ApplicationId)
	if err != nil {
		return err
	}

	err = d.Set("api_key", response.ApiKey)

	if err != nil {
		return err
	}

	return applicationRead(d, m)
}

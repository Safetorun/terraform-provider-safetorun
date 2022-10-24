package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	AppConfiguration = "application_configuration_filepath"
)

func applicationConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Create: applicationConfigurationCreate,
		Read:   applicationConfigurationRead,
		Update: applicationConfigurationCreate,
		Delete: applicationConfigurationDelete,

		Schema: map[string]*schema.Schema{
			AppConfiguration: {
				Type:     schema.TypeString,
				Required: true,
			},
			OrganisationId: {
				Type:     schema.TypeString,
				Required: true,
			},
			ApplicationId: {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func applicationConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func applicationConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func applicationConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	organisationId := d.Get(OrganisationId).(string)
	applicationId := d.Get(ApplicationId).(string)
	applicationConfiguration := d.Get(AppConfiguration).(string)

	response, err := m.(SafeToRunProvider).Client.UploadApplicationConfiguration(safetorun.UploadApplicationConfiguration{
		ConfigurationFilename: applicationConfiguration,
		ApplicationId:         applicationId,
		OrganisationId:        organisationId,
	})

	if err != nil {
		return err
	}

	d.SetId(response.OrganisationId)

	if err != nil {
		return err
	}

	return resourceServerRead(d, m)
}

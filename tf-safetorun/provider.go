package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type SafeToRunProvider struct {
	Client safetorun.Client
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: func(data *schema.ResourceData) (interface{}, error) {
			return SafeToRunProvider{Client: safetorun.New(data.Get("token").(string))}, nil
		},
		Schema: map[string]*schema.Schema{
			"token": accessToken(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"safetorun_organisation":              organisationResource(),
			"safetorun_application":               applicationResource(),
			"safetorun_application_configuration": applicationConfigurationResource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"safetorun_organisation": organisationData(),
		},
	}
}

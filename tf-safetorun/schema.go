package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func accessToken() *schema.Schema {
	return &schema.Schema{
		Required:  true,
		Type:      schema.TypeString,
		Sensitive: true,
	}
}

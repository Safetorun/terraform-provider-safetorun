package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "safetorun_test")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("safetorun"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("safetorun_test")),
	})

	app.Synth()
}

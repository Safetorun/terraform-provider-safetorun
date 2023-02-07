# Quickstart

```terraform

variable "safetorun_token" {
  type = string

}
provider "safetorun" {
  token = var.safetorun_token
}

resource "safetorun_organisation" "myorg" {
  organisation_id   = "test_api_infrastructure"
  organisation_name = "Test organisation"
}

resource "safetorun_application" "app" {
  organisation_id  = safetorun_organisation.myorg.id
  application_name = "Test application name"
}
```
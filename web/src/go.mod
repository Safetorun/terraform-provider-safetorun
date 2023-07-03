module github.com/Safetorun/terraform-provider-safetorun/web

go 1.19

replace github.com/Safetorun/safe_to_run_admin_api/safetorun => ../../safetorun

require (
	github.com/Safetorun/safe_to_run_admin_api/safetorun v0.0.0-20230501191310-6b05f63aa17e
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/Khan/genqlient v0.6.0 // indirect
	github.com/amplitude/analytics-go v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

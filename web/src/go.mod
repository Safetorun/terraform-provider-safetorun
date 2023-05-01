module github.com/Safetorun/terraform-provider-safetorun/web

go 1.19

replace github.com/Safetorun/safe_to_run_admin_api/safetorun => ../../safetorun

require (
	github.com/Safetorun/safe_to_run_admin_api/safetorun v0.0.0-20230313095708-0ff25f81d82d
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/Khan/genqlient v0.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

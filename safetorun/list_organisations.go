package safetorun

import (
	"context"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
)

func (client Client) ListOrganisations() (*ListOrganisationsListOrganisationsOrganisationList, error) {
	ctx := context.Background()

	client.logListOrganisationAnalytics()
	response, err := ListOrganisations(ctx, client.GqlClient)

	if err != nil {
		return nil, err
	}

	resp := response.GetListOrganisations()

	return &resp, nil
}

func (client Client) logListOrganisationAnalytics() {
	createOrg := ampli.ListOrgs.Builder().Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

func (client Client) DeleteOrganisation(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to Delete organisation for ID %s", organisationId))
	client.logDeleteOrgAnalytics(organisationId)

	response, err := DeleteOrganisation(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetDeleteOrganisation()
	return &status, nil
}

func (client Client) DeleteOrganisationAndWait(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	return PerformActionAndWait(client, organisationId, organisationId, client.DeleteOrganisation)
}

func (client Client) logDeleteOrgAnalytics(organisationId string) {
	createOrg := ampli.DeleteOrganisation.Builder().OrganisationId(organisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

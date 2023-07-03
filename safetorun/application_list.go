package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

func (client Client) ListApplications(organisationId string) (*GetApplicationsListApplicationsApplicationList, error) {
	ctx := context.Background()

	client.logListsAppAnalytics(organisationId)
	log.Println(fmt.Sprintf("Going to list application for the organisation %s", organisationId))
	response, err := GetApplications(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.ListApplications
	return &status, err
}

func (client Client) logListsAppAnalytics(organisationId string) {
	createOrg := ampli.ListApplications.Builder().OrganisationId(organisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

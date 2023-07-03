package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

type DeleteApplicationRequest struct {
	OrganisationId string
	ApplicationId  string
}

type DeleteApplicationResp struct {
	ApplicationId string
}

func (client Client) DeleteApplication(request DeleteApplicationRequest) (*DeleteApplicationResp, error) {

	ctx := context.Background()

	client.logDeleteAppAnalytics(request.OrganisationId, request.ApplicationId)

	log.Println(fmt.Sprintf("Going to delete applciation with organisation ID: %s and application ID: %s", request.OrganisationId, request.ApplicationId))
	response, err := DeleteApplication(ctx, client.GqlClient, request.OrganisationId, request.ApplicationId)

	if err != nil {
		return nil, err
	}

	status := response.DeleteApplication

	return &DeleteApplicationResp{
		ApplicationId: status.ApplicationId,
	}, err
}

func (client Client) DeleteApplicationAndWait(request DeleteApplicationRequest) (*DeleteApplicationResp, error) {
	return PerformActionAndWait(client, request, request.OrganisationId, client.DeleteApplication)
}

func (client Client) logDeleteAppAnalytics(organisationId string, applicationId string) {
	createOrg := ampli.DeleteApplication.Builder().ApplicationId(applicationId).OrganisationId(organisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

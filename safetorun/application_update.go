package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

type UpdateApplicationRequest struct {
	OrganisationId  string
	ApplicationId   string
	ApplicationName string
}

func (client Client) UpdateApplication(request UpdateApplicationRequest) (*UpdateApplicationUpdateApplicationCreateApplicationResponse, error) {
	ctx := context.Background()
	client.logUpdateAppAnalytics(request.ApplicationId, request.OrganisationId)
	log.Println(fmt.Sprintf("Going to update application for the organisation %s and appId %s to have new app name %s", request.OrganisationId, request.ApplicationId, request.ApplicationName))
	response, err := UpdateApplication(ctx, client.GqlClient, request.OrganisationId, request.ApplicationId, request.ApplicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetUpdateApplication()
	return &status, err
}

func (client Client) UpdateApplicationAndWait(request UpdateApplicationRequest) (*UpdateApplicationUpdateApplicationCreateApplicationResponse, error) {
	return PerformActionAndWait(client, request, request.OrganisationId, client.UpdateApplication)
}

func (client Client) logUpdateAppAnalytics(organisationId string, applicationId string) {
	createOrg := ampli.UpdateApplication.Builder().ApplicationId(applicationId).OrganisationId(organisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

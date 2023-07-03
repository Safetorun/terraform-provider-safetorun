package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

type CreateApplicationRequest struct {
	OrganisationId  string
	ApplicationName string
}

type CreateApplicationResp struct {
	ApiKey        string
	ApplicationId string
}

func (client Client) CreateApplication(request CreateApplicationRequest) (*CreateApplicationResp, error) {

	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to create application inside org ID %s with name %s", request.OrganisationId, request.ApplicationName))

	client.logCreateAppAnalytics(request)
	response, err := CreateApplication(ctx, client.GqlClient, request.OrganisationId, request.ApplicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetCreateApplication()

	return &CreateApplicationResp{
		ApplicationId: status.ApplicationId,
	}, err
}

func (client Client) CreateApplicationAndWait(request CreateApplicationRequest) (*CreateApplicationResp, error) {
	return PerformActionAndWait(client, request, request.OrganisationId, client.CreateApplication)
}

func (client Client) logCreateAppAnalytics(request CreateApplicationRequest) {
	createOrg := ampli.CreateApplication.Builder().OrganisationId(request.OrganisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

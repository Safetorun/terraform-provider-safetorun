package safetorun

import (
	"context"
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"log"
)

type CreateOrganisationRequest struct {
	OrganisationName string
	OrganisationId   string
}

type CreateOrganisationResp struct {
	Status         int
	OrganisationId string
}

func (client Client) CreateOrganisationAndWait(request CreateOrganisationRequest) (*CreateOrganisationResp, error) {
	return PerformActionAndWait(client, request, request.OrganisationId, client.CreateOrganisation)
}

func (client Client) CreateOrganisation(request CreateOrganisationRequest) (*CreateOrganisationResp, error) {

	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to create organisation for ID %s with name %s", request.OrganisationId, request.OrganisationName))

	client.logCreateOrganisationAnalytics(request)

	response, err := CreateOrganisation(ctx, client.GqlClient, request.OrganisationId, request.OrganisationName)

	if err != nil {
		return nil, err
	}

	status := response.GetCreateOrganisation()

	return &CreateOrganisationResp{
		Status:         status.Status,
		OrganisationId: status.OrganisationId,
	}, err
}

func (client Client) logCreateOrganisationAnalytics(request CreateOrganisationRequest) {
	createOrg := ampli.CreateOrganisation.Builder().OrganisationId(request.OrganisationId).Build()
	ampli.Instance.Track(client.UserId, createOrg)
}

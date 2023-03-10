package safetorun

import (
	"context"
	"fmt"
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

func (client Client) ListOrganisations() (*ListOrganisationsListOrganisationsOrganisationList, error) {
	ctx := context.Background()
	response, err := ListOrganisations(ctx, client.GqlClient)

	if err != nil {
		return nil, err
	}

	resp := response.GetListOrganisations()

	return &resp, nil
}

func (client Client) CreateOrganisation(request CreateOrganisationRequest) (*CreateOrganisationResp, error) {

	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to create organisation for ID %s with name %s", request.OrganisationId, request.OrganisationName))
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

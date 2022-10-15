package safetorun

import (
	"context"
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

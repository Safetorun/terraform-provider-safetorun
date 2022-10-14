package safetorun

import (
	"context"
)

type CreateOrganisationRequest struct {
	OrganisationName string
	OrganisationId   string
	AdminUser        string
}

type CreateOrganisationResp struct {
	Status         int
	OrganisationId string
}

func (client Client) CreateOrganisation(request CreateOrganisationRequest) (*CreateOrganisationResp, error) {

	ctx := context.Background()

	response, err := CreateOrganisation(ctx, client.GqlClient, request.AdminUser, request.OrganisationId, request.OrganisationName)

	if err != nil {
		return nil, err
	}

	status := response.GetCreateOrganisation()

	return &CreateOrganisationResp{
		Status:         status.Status,
		OrganisationId: status.OrganisationId,
	}, err
}

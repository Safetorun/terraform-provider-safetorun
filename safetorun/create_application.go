package safetorun

import (
	"context"
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

	response, err := CreateApplication(ctx, client.GqlClient, request.OrganisationId, request.ApplicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetCreateApplication()

	return &CreateApplicationResp{
		ApplicationId: status.ApplicationId,
	}, err
}

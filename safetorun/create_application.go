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

func (c Client) CreateApplication(request CreateApplicationRequest) (*CreateApplicationResp, error) {

	ctx := context.Background()

	response, err := CreateApplication(ctx, c.GqlClient, request.OrganisationId, request.ApplicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetCreateApplication()

	return &CreateApplicationResp{
		ApplicationId: status.ApplicationId,
	}, err
}

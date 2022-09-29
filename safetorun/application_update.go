package safetorun

import "context"

func (c Client) UpdateApplication(organisationId string, applicationId string, applicationName string) (*UpdateApplicationUpdateApplicationCreateApplicationResponse, error) {
	ctx := context.Background()

	response, err := UpdateApplication(ctx, c.GqlClient, organisationId, applicationId, applicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetUpdateApplication()
	return &status, err
}

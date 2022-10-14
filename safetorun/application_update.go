package safetorun

import "context"

func (client Client) UpdateApplication(organisationId string, applicationId string, applicationName string) (*UpdateApplicationUpdateApplicationCreateApplicationResponse, error) {
	ctx := context.Background()

	response, err := UpdateApplication(ctx, client.GqlClient, organisationId, applicationId, applicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetUpdateApplication()
	return &status, err
}

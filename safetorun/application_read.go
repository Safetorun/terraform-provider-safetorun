package safetorun

import "context"

func (client Client) QueryApplication(organisationId string, applicationId string) (*GetApplicationGetApplication, error) {
	ctx := context.Background()

	response, err := GetApplication(ctx, client.GqlClient, organisationId, applicationId)

	if err != nil {
		return nil, err
	}

	status := response.GetGetApplication()
	return &status, err
}

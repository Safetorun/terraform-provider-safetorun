package safetorun

import "context"

func (c Client) QueryApplication(organisationId string, applicationId string) (*GetApplicationGetApplication, error) {
	ctx := context.Background()

	response, err := GetApplication(ctx, c.GqlClient, organisationId, applicationId)

	if err != nil {
		return nil, err
	}

	status := response.GetGetApplication()
	return &status, err
}

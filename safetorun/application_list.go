package safetorun

import "context"

func (client Client) ListApplications(organisationId string) (*GetApplicationsListApplicationsApplicationList, error) {
	ctx := context.Background()

	response, err := GetApplications(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.ListApplications
	return &status, err
}

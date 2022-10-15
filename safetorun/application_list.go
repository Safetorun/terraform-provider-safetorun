package safetorun

import (
	"context"
	"fmt"
	"log"
)

func (client Client) ListApplications(organisationId string) (*GetApplicationsListApplicationsApplicationList, error) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to list application for the organisation %s", organisationId))
	response, err := GetApplications(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.ListApplications
	return &status, err
}

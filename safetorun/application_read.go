package safetorun

import (
	"context"
	"fmt"
	"log"
)

func (client Client) QueryApplication(organisationId string, applicationId string) (*GetApplicationGetApplication, error) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to look for application for the organisation %s and appId %s", organisationId, applicationId))
	response, err := GetApplication(ctx, client.GqlClient, organisationId, applicationId)

	if err != nil {
		return nil, err
	}

	status := response.GetGetApplication()
	return &status, err
}

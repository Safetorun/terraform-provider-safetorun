package safetorun

import (
	"context"
	"fmt"
	"log"
)

func (client Client) UpdateApplication(organisationId string, applicationId string, applicationName string) (*UpdateApplicationUpdateApplicationCreateApplicationResponse, error) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to update application for the organisation %s and appId %s to have new app name %s", organisationId, applicationId, applicationName))
	response, err := UpdateApplication(ctx, client.GqlClient, organisationId, applicationId, applicationName)

	if err != nil {
		return nil, err
	}

	status := response.GetUpdateApplication()
	return &status, err
}

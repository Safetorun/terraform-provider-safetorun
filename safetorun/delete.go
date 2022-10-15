package safetorun

import (
	"context"
	"fmt"
	"log"
)

func (client Client) DeleteOrganisation(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to Delete organisation for ID %s", organisationId))

	response, err := DeleteOrganisation(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetDeleteOrganisation()
	return &status, nil
}

func (client Client) DeleteOrganisationAndWait(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	return PerformActionAndWait(client, organisationId, organisationId, client.DeleteOrganisation)
}

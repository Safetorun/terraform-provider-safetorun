package safetorun

import "context"

func (client Client) DeleteOrganisation(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	ctx := context.Background()

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

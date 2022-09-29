package safetorun

import "context"

func (c Client) DeleteOrganisation(organisationId string) (*DeleteOrganisationDeleteOrganisationOrganisationStatus, error) {
	ctx := context.Background()

	response, err := DeleteOrganisation(ctx, c.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetDeleteOrganisation()
	return &status, nil
}

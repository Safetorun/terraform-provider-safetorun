package safetorun

import "context"

func (c Client) DeleteOrganisation(organisationId string) (*DeleteDeleteOrganisationOrganisationStatus, error) {
	ctx := context.Background()

	response, err := Delete(ctx, c.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetDeleteOrganisation()
	return &status, nil
}

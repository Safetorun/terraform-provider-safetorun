package safetorun

import "context"

func (client Client) ListOrganisations() (*ListOrganisationsListOrganisationsOrganisationList, error) {
	ctx := context.Background()
	response, err := ListOrganisations(ctx, client.GqlClient)

	if err != nil {
		return nil, err
	}

	resp := response.GetListOrganisations()

	return &resp, nil
}

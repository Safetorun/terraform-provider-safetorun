package safetorun

import (
	"context"
)

type OrgStatusResult struct {
	GetOrganisationStatus OrganisationStatus `json:"getOrganisationStatus"`
}
type OrganisationStatus struct {
	OrganisationName string
	Status           CreateStatus
}

type CreateStatus int

const (
	CreateInProgress      = iota
	InfrastructureCreated = iota
	ErrorDestroying       = iota
	DeleteComplete        = iota
	AlreadyExists         = iota
)

func (c Client) QueryStatus(organisationId string) (*GetForOrganisationIdGetOrganisationStatus, error) {
	ctx := context.Background()

	response, err := GetForOrganisationId(ctx, c.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetGetOrganisationStatus()
	return &status, err
}

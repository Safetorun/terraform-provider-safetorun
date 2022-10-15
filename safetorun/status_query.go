package safetorun

import (
	"context"
	"errors"
	"log"
	"time"
)

type OrgStatusResult struct {
	GetOrganisationStatus OrganisationStatus `json:"getOrganisationStatus"`
}
type OrganisationStatus struct {
	OrganisationName string
	Status           Status
}

type Status int

const (
	CreateInProgress Status = iota
	EndedWithError          = iota
	EndedWithSuccess        = iota
	Noop                    = iota
)

func (client Client) QueryStatus(organisationId string) (*GetForOrganisationIdGetOrganisationStatus, error) {
	ctx := context.Background()

	response, err := GetForOrganisationId(ctx, client.GqlClient, organisationId)

	if err != nil {
		return nil, err
	}

	status := response.GetGetOrganisationStatus()
	return &status, err
}

func PerformActionAndWait[T, U any](client Client, u T, organisationId string, f func(T) (*U, error)) (*U, error) {
	response, err := f(u)
	if err != nil {
		return nil, err
	}

	err = client.WaitForCompletion(organisationId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client Client) WaitForCompletion(organisationId string) error {
	for {
		re, err := client.RetrieveLastEventForLinkId(organisationId)

		if err != nil {
			log.Fatal(err)
		}

		switch re.Status {
		case int(CreateInProgress):
			time.Sleep(time.Second)
			break
		case EndedWithSuccess:
			return nil
		case EndedWithError:
			return errors.New("failed to complete")
		}
	}
}

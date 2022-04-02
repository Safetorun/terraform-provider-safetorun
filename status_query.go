package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
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
	CreateInProgress      CreateStatus = iota
	InfrastructureCreated              = iota
	ErrorDestroying                    = iota
	DeleteComplete                     = iota
	AlreadyExists                      = iota
	Noop                               = iota
)

func QueryStatus(organisationName string, authToken string) (*OrganisationStatus, error) {
	query := fmt.Sprintf(`query MyQuery {
		  getOrganisationStatus(organisationName: "%s") {
			OrganisationName
			Status
		  }
		}
		`, organisationName)

	req := graphql.NewRequest(query)
	url := "https://ulhdaocpgfewxmt7xxf55l6mzm.appsync-api.eu-west-1.amazonaws.com/graphql"
	client := graphql.NewClient(url)
	client.Log = func(s string) { log.Println(s) }
	req.Header.Set("Authorization", authToken)

	var status OrgStatusResult
	err := client.Run(context.Background(), req, &status)

	if err != nil {
		return nil, err
	}

	return &status.GetOrganisationStatus, nil

}

package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var organisationName string
	var authToken string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "org_name",
				Usage:       "Organisation name to create",
				Aliases:     []string{"o"},
				Destination: &organisationName,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "auth_token",
				Usage:       "Token for authentication",
				Aliases:     []string{"a"},
				Destination: &authToken,
				Required:    true,
			},
		},
		Name:  "create",
		Usage: "Create a new application on safe to run",
		Action: func(c *cli.Context) error {
			re, err := QueryStatus(organisationName, authToken)

			if err != nil {
				log.Fatal(err)
			}

			println(fmt.Sprintf("%+v", re))
			response, err := CreateOrganisation(CreateOrganisationRequest{
				OrganisationName: organisationName,
				OrganisationId:   "",
				AdminUser:        "",
			}, authToken)

			if err != nil {
				log.Fatal(err)
			}

			println(fmt.Sprintf("%+v", *response))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

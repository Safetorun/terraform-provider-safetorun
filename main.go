package main

import (
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var organisationName string
	var applicationName string
	var applicationId string
	var authToken string
	var organisationId string

	app := &cli.App{
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "auth_token",
				Usage:       "Token for authentication",
				Aliases:     []string{"a"},
				Destination: &authToken,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "organisation_id",
				Destination: &organisationId,
				Required:    true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "list-apps",
				Usage: "List applications",
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					event, err := client.ListApplications(organisationId)

					if err != nil {
						log.Fatal(err)
						return err
					}

					log.Println(fmt.Sprintf("%+v", event))
					return nil
				},
			},
			{
				Name:  "latest_event",
				Usage: "Latest event",
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					event, err := client.RetrieveLastEventForLinkId(organisationId)

					if err != nil {
						log.Fatal(err)
						return err
					}

					log.Println(fmt.Sprintf("%+v", event))
					return nil
				},
			},
			{
				Name:  "delete-org",
				Usage: "Delete an organisation from safe to run",
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.DeleteOrganisation(organisationId)

					if err != nil {
						log.Fatal(err)
						return err
					}

					return client.WaitForCompletion(organisationId)
				},
			},
			{
				Name:  "create-app",
				Usage: "Create a new application on safe to run",
				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:        "application_name",
						Usage:       "Application name to create",
						Aliases:     []string{"a"},
						Destination: &applicationName,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.CreateApplication(safetorun.CreateApplicationRequest{
						OrganisationId:  organisationId,
						ApplicationName: applicationName,
					})

					if err != nil {
						log.Fatal(err)
					}

					return client.WaitForCompletion(organisationId)
				},
			},
			{
				Name:  "delete-app",
				Usage: "delete a new application on safe to run",
				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:        "application_id",
						Usage:       "Application name to create",
						Aliases:     []string{"aid"},
						Destination: &applicationId,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.DeleteApplication(safetorun.DeleteApplicationRequest{
						OrganisationId: organisationId,
						ApplicationId:  applicationId,
					})

					if err != nil {
						log.Fatal(err)
					}

					return err
				},
			},
			{
				Name:  "create-org",
				Usage: "Create a new organisation on safe to run",
				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:        "org_name",
						Usage:       "Organisation name to create",
						Aliases:     []string{"o"},
						Destination: &organisationName,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.CreateOrganisation(safetorun.CreateOrganisationRequest{
						OrganisationName: organisationName,
						OrganisationId:   organisationId,
					})

					if err != nil {
						log.Fatal(err)
					}

					return client.WaitForCompletion(organisationId)
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

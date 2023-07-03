package main

import (
	"fmt"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
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
	var applicationFile cli.Path

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
		},
		Commands: []*cli.Command{
			{
				Name:  "list-apps",
				Usage: "List applications",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
				},
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
				Name:  "get-app",
				Usage: "Get application",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "application_id",
						Usage:       "Application name to create",
						Aliases:     []string{"aid"},
						Destination: &applicationId,
						Required:    true,
					},
				},
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					event, err := client.QueryApplication(organisationId, applicationId)

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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
				},
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
				Name: "delete-org",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
				},
				Usage: "Delete an organisation from safe to run",
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.DeleteOrganisationAndWait(organisationId)
					return err
				},
			},
			{
				Name:  "upload-config",
				Usage: "Upload application config",
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.UploadApplicationConfiguration(safetorun.UploadApplicationConfiguration{
						OrganisationId:        organisationId,
						ApplicationId:         applicationId,
						ConfigurationFilename: applicationFile,
					})
					return err
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},

					&cli.StringFlag{
						Name:        "application_id",
						Usage:       "Application name to create",
						Aliases:     []string{"aid"},
						Destination: &applicationId,
						Required:    true,
					},
					&cli.PathFlag{
						Destination: &applicationFile,
						Name:        "application_config",
						Aliases:     []string{"c"},
						Required:    true,
					},
				},
			},
			{
				Name:  "create-app",
				Usage: "Create a new application on safe to run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
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
					_, err := client.CreateApplicationAndWait(safetorun.CreateApplicationRequest{
						OrganisationId:  organisationId,
						ApplicationName: applicationName,
					})

					return err
				},
			},
			{
				Name:  "delete-app",
				Usage: "delete a new application on safe to run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
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
					_, err := client.DeleteApplicationAndWait(safetorun.DeleteApplicationRequest{
						OrganisationId: organisationId,
						ApplicationId:  applicationId,
					})
					return err
				},
			},
			{
				Name:  "create-org",
				Usage: "Create a new organisation on safe to run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "organisation_id",
						Destination: &organisationId,
						Required:    true,
					},
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
					_, err := client.CreateOrganisationAndWait(safetorun.CreateOrganisationRequest{
						OrganisationName: organisationName,
						OrganisationId:   organisationId,
					})

					return err
				},
			},
			{
				Name:  "list-orgs",
				Usage: "List organisations for this user",
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					organisations, err := client.ListOrganisations()

					if err != nil {
						return err
					}

					for _, organisation := range organisations.Items {
						log.Println(fmt.Sprintf("Orgsanisation Id: %+v", organisation))
					}
					return err
				},
			},
		},
	}
	err := app.Run(os.Args)

	ampli.Instance.Flush()

	if err != nil {
		log.Fatal(err)
	}

}

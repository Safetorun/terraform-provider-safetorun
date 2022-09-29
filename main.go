package main

import (
	"github.com/Safetorun/safe_to_run_admin_api/safetorun"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	var organisationName string
	var applicationName string
	var applicationId string
	var authToken string
	var organisationId string
	var adminEmail string

	app := &cli.App{
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
				Name:  "delete",
				Usage: "Delete an organisation from safe to run",
				Action: func(context *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.DeleteOrganisation(organisationId)

					if err != nil {
						log.Fatal(err)
						return err
					}

					return waitForStatus(client, organisationId)
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

					return err
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
				Name:  "create",
				Usage: "Create a new organisation on safe to run",
				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:        "org_name",
						Usage:       "Organisation name to create",
						Aliases:     []string{"o"},
						Destination: &organisationName,
						Required:    true,
					},

					&cli.StringFlag{
						Name:        "admin_email",
						Usage:       "Admin email",
						Destination: &adminEmail,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					client := safetorun.New(authToken)
					_, err := client.CreateOrganisation(safetorun.CreateOrganisationRequest{
						OrganisationName: organisationName,
						OrganisationId:   organisationId,
						AdminUser:        adminEmail,
					})

					if err != nil {
						log.Fatal(err)
					}

					return waitForStatus(client, organisationId)
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func waitForStatus(client safetorun.Client, organisationId string) error {
	for {
		re, err := client.QueryStatus(organisationId)

		if err != nil {
			log.Fatal(err)
		}

		switch re.Status {
		case safetorun.CreateInProgress:
			time.Sleep(time.Second)
			break
		case safetorun.InfrastructureCreated:
			println("Create complete")
			return nil

		case safetorun.ErrorDestroying:
			println("Something went wrong, destroying.")
			time.Sleep(time.Second)
			break
		case safetorun.DeleteComplete:
			println("Delete complete.")
			return nil
		case safetorun.AlreadyExists:
			println("Org already exists")
			return nil
		}
	}
}

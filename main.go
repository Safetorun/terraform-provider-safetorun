package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	var organisationName string
	var authToken string
	var organisationId string
	var adminEmail string

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
			&cli.StringFlag{
				Name:        "admin_email",
				Usage:       "Admin email",
				Destination: &adminEmail,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "organisation_id",
				Destination: &organisationId,
				Required:    true,
			},
		},
		Name:  "create",
		Usage: "Create a new application on safe to run",
		Action: func(c *cli.Context) error {

			client := New(authToken)
			_, err := client.CreateOrganisation(CreateOrganisationRequest{
				OrganisationName: organisationName,
				OrganisationId:   organisationId,
				AdminUser:        adminEmail,
			})

			if err != nil {
				log.Fatal(err)
			}

			for {
				re, err := client.QueryStatus(organisationName)

				if err != nil {
					log.Fatal(err)
				}

				switch re.Status {
				case CreateInProgress:
					time.Sleep(time.Second)
					break
				case InfrastructureCreated:
					println("Create complete")
					return nil

				case ErrorDestroying:
					println("Something went wrong, destroying.")
					time.Sleep(time.Second)
					break
				case DeleteComplete:
					println("Delete complete.")
					return nil
				case AlreadyExists:
					println("Org already exists")
					return nil
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

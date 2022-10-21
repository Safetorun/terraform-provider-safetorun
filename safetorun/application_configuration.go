package safetorun

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type UploadApplicationConfiguration struct {
	ConfigurationFilename string
	ApplicationId         string
	OrganisationId        string
}

func (client Client) UploadApplicationConfiguration(request UploadApplicationConfiguration) (*UploadApplicationConfiguration, error) {

	ctx := context.Background()

	log.Println(fmt.Sprintf("Going to upload new configuration for org %s and app %s", request.OrganisationId, request.ApplicationId))

	uploadUrl, err := UploadUrl(ctx, client.GqlClient, request.OrganisationId, request.ApplicationId)

	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("Have URL %s. Going to upload file now", uploadUrl.GetUploadUrl.Url))
	fileBytes, err := os.ReadFile(request.ConfigurationFilename)

	if err != nil {
		log.Fatal(err)
	}

	r, err := http.NewRequest(http.MethodPut, uploadUrl.GetGetUploadUrl().Url, bytes.NewReader(fileBytes))

	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{}
	re, err := c.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	if re.StatusCode > 299 || re.StatusCode < 200 {
		return nil, errors.New(fmt.Sprintf("Error response. %s", re.Status))
	}

	return &request, nil
}

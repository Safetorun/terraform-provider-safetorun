package safetorun

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	URL = "https://c9w5whh2jc.execute-api.eu-west-1.amazonaws.com/Prod"
)

type CreateOrganisationRequest struct {
	OrganisationName string `json:"organisation_name"`
	OrganisationId   string `json:"organisation_id"`
	AdminUser        string `json:"admin_user"`
}

type CreateOrganisationResponse struct {
	Message        string `json:"message"`
	OrganisationId string `json:"organisation_id"`
}

func (c Client) CreateOrganisation(request CreateOrganisationRequest) (*CreateOrganisationResponse, error) {
	body, err := json.Marshal(request)

	writer := bytes.NewBuffer(body)

	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/organisation/create", URL), writer)

	r.Header.Set("Bearer", c.AuthToken)

	client := http.Client{}

	response, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, errors.New("unexpected response, error and response both nil")
	}

	if response.StatusCode == 401 {
		return nil, errors.New("failed to authenticate")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	responseData := CreateOrganisationResponse{}
	err = json.Unmarshal(responseBody, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

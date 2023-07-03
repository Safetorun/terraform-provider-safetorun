// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package safetorun

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetUserIdResponse is returned by GetUserId on success.
type GetUserIdResponse struct {
	Whoami GetUserIdWhoamiWhoamiResponse `json:"whoami"`
}

// GetWhoami returns GetUserIdResponse.Whoami, and is useful for accessing the field via an interface.
func (v *GetUserIdResponse) GetWhoami() GetUserIdWhoamiWhoamiResponse { return v.Whoami }

// GetUserIdWhoamiWhoamiResponse includes the requested fields of the GraphQL type WhoamiResponse.
type GetUserIdWhoamiWhoamiResponse struct {
	UserId string `json:"UserId"`
}

// GetUserId returns GetUserIdWhoamiWhoamiResponse.UserId, and is useful for accessing the field via an interface.
func (v *GetUserIdWhoamiWhoamiResponse) GetUserId() string { return v.UserId }

// __GetUserIdInput is used internally by genqlient
type __GetUserIdInput struct {
	AuthorisationToken string `json:"authorisationToken"`
}

// GetAuthorisationToken returns __GetUserIdInput.AuthorisationToken, and is useful for accessing the field via an interface.
func (v *__GetUserIdInput) GetAuthorisationToken() string { return v.AuthorisationToken }

// The query or mutation executed by GetUserId.
const GetUserId_Operation = `
query GetUserId ($authorisationToken: String!) {
	whoami(request: {authorizationToken:$authorisationToken}) {
		UserId
	}
}
`

func GetUserId(
	ctx context.Context,
	client graphql.Client,
	authorisationToken string,
) (*GetUserIdResponse, error) {
	req := &graphql.Request{
		OpName: "GetUserId",
		Query:  GetUserId_Operation,
		Variables: &__GetUserIdInput{
			AuthorisationToken: authorisationToken,
		},
	}
	var err error

	var data GetUserIdResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

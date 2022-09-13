package safetorun

import (
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

type Client struct {
	GqlClient graphql.Client
}

type myTransport struct {
	Header string
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.Header))
	return http.DefaultTransport.RoundTrip(req)
}

func AuthenticatedClient(header string) *http.Client {
	return &http.Client{Transport: &myTransport{
		Header: header,
	}}
}

func GqlClient(header string) graphql.Client {
	httpClient := AuthenticatedClient(header)
	return graphql.NewClient("https://j4tqmk277rccrcliofgynrp774.appsync-api.eu-west-1.amazonaws.com/graphql", httpClient)
}

func New(authToken string) Client {
	return Client{
		GqlClient(authToken),
	}
}

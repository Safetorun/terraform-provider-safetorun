package safetorun

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/ampli"
	"github.com/Safetorun/safe_to_run_admin_api/safetorun/logger"
	"github.com/amplitude/analytics-go/amplitude"
	"net/http"
)

type Client struct {
	GqlClient graphql.Client
	UserId    string
}

type myTransport struct {
	Header string
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", t.Header)
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

func AuthClient(header string) graphql.Client {
	httpClient := AuthenticatedClient(header)
	return graphql.NewClient("https://26btv5v46jb3lbpqzhxccqzugi.appsync-api.eu-west-1.amazonaws.com/graphql", httpClient)
}

func New(authToken string) Client {
	gqlClient := GqlClient(authToken)
	authClient := AuthClient(authToken)

	id, err := GetUserId(context.Background(), authClient, authToken)

	if err != nil {
		panic(err)
	}

	configureAmplitude(id)

	return Client{
		GqlClient: gqlClient,
		UserId:    id.Whoami.UserId,
	}
}

func configureAmplitude(id *GetUserIdResponse) {
	ampli.Instance.Load(ampli.LoadOptions{
		Environment: ampli.EnvironmentProd,
		Client: ampli.LoadClientOptions{
			Configuration: amplitude.Config{
				Logger: logger.MockLogger{},
			},
		},
	})

	ampli.Instance.Identify(id.Whoami.UserId)
}

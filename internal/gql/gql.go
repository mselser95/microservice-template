package gql

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

// GQLClient is an interface for making GraphQL requests
type GQLClient interface {
	MakeRequest(
		ctx context.Context,
		req *graphql.Request,
		resp *graphql.Response,
	) error
}

// NewGQLClient creates a new GQLClient
func NewGQLClient(endPoint, adminSecret, adminRole string) GQLClient {
	hc := http.DefaultClient
	hc.Transport = &authTransport{
		next:        http.DefaultTransport,
		adminSecret: adminSecret,
		adminRole:   adminRole,
	}
	return graphql.NewClient(endPoint, hc)
}

type authTransport struct {
	next        http.RoundTripper
	adminSecret string
	adminRole   string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-hasura-admin-secret", t.adminSecret)
	req.Header.Set("x-hasura-role", t.adminRole)
	return t.next.RoundTrip(req)
}

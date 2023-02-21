package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/pluralsh/oauth-playground/api-server/format"
	"github.com/pluralsh/oauth-playground/api-server/graph/model"
)

// ListOAuth2Clients is the resolver for the listOAuth2Clients field.
func (r *queryResolver) ListOAuth2Clients(ctx context.Context) ([]*model.OAuth2Client, error) {
	log := r.C.Log.WithName("ListOAuth2Clients")
	clients, resp, err := r.C.HydraClient.OAuth2Api.ListOAuth2Clients(ctx).Execute()
	if err != nil || resp.StatusCode != 200 {
		log.Error(err, "failed to list oauth2 clients")
		return nil, fmt.Errorf("failed to list oauth2 clients: %w", err)
	}
	var output []*model.OAuth2Client
	for _, client := range clients {
		output = append(output, format.HydraOAuth2ClientToGraphQL(client))
	}
	return output, nil
}

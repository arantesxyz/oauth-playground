package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/pluralsh/oauth-playground/api-server/clients"
	"github.com/pluralsh/oauth-playground/api-server/graph/generated"
	"github.com/pluralsh/oauth-playground/api-server/graph/model"
	"github.com/pluralsh/oauth-playground/api-server/utils"
)

// Users is the resolver for the users field.
func (r *loginBindingsResolver) Users(ctx context.Context, obj *model.LoginBindings) ([]*model.User, error) {
	return r.C.GetOAuth2ClientUserLoginBindings(ctx, obj)
}

// Groups is the resolver for the groups field.
func (r *loginBindingsResolver) Groups(ctx context.Context, obj *model.LoginBindings) ([]*model.Group, error) {
	return r.C.GetOAuth2ClientGroupLoginBindings(ctx, obj)
}

// CreateOAuth2Client is the resolver for the createOAuth2Client field.
func (r *mutationResolver) CreateOAuth2Client(ctx context.Context, allowedCorsOrigins []string, audience []string, authorizationCodeGrantAccessTokenLifespan *string, authorizationCodeGrantIDTokenLifespan *string, authorizationCodeGrantRefreshTokenLifespan *string, backChannelLogoutSessionRequired *bool, backChannelLogoutURI *string, clientCredentialsGrantAccessTokenLifespan *string, clientName *string, clientSecret *string, clientSecretExpiresAt *int64, clientURI *string, contacts []string, frontchannelLogoutSessionRequired *bool, frontchannelLogoutURI *string, grantTypes []string, implicitGrantAccessTokenLifespan *string, implicitGrantIDTokenLifespan *string, jwks map[string]interface{}, jwksURI *string, jwtBearerGrantAccessTokenLifespan *string, logoURI *string, metadata map[string]interface{}, policyURI *string, postLogoutRedirectUris []string, redirectUris []string, responseTypes []string, scope *string, sectorIdentifierURI *string, subjectType *string, tokenEndpointAuthMethod *string, tokenEndpointAuthSigningAlgorithm *string, tosURI *string, userinfoSignedResponseAlgorithm *string, loginBindings *model.LoginBindingsInput) (*model.OAuth2Client, error) {
	if clientSecret == nil || *clientSecret == "" {
		secret := utils.RandStringBytesMaskImprSrcUnsafe(32)
		clientSecret = &secret
	}
	return r.C.CreateOAuth2Client(ctx, clients.HydraOperationCreate, allowedCorsOrigins, audience, authorizationCodeGrantAccessTokenLifespan, authorizationCodeGrantIDTokenLifespan, authorizationCodeGrantRefreshTokenLifespan, backChannelLogoutSessionRequired, backChannelLogoutURI, clientCredentialsGrantAccessTokenLifespan, nil, clientName, clientSecret, clientSecretExpiresAt, clientURI, contacts, frontchannelLogoutSessionRequired, frontchannelLogoutURI, grantTypes, implicitGrantAccessTokenLifespan, implicitGrantIDTokenLifespan, jwks, jwksURI, jwtBearerGrantAccessTokenLifespan, logoURI, metadata, policyURI, postLogoutRedirectUris, redirectUris, responseTypes, scope, sectorIdentifierURI, subjectType, tokenEndpointAuthMethod, tokenEndpointAuthSigningAlgorithm, tosURI, userinfoSignedResponseAlgorithm, loginBindings)
}

// UpdateOAuth2Client is the resolver for the updateOAuth2Client field.
func (r *mutationResolver) UpdateOAuth2Client(ctx context.Context, allowedCorsOrigins []string, audience []string, authorizationCodeGrantAccessTokenLifespan *string, authorizationCodeGrantIDTokenLifespan *string, authorizationCodeGrantRefreshTokenLifespan *string, backChannelLogoutSessionRequired *bool, backChannelLogoutURI *string, clientCredentialsGrantAccessTokenLifespan *string, clientID string, clientName *string, clientSecret *string, clientSecretExpiresAt *int64, clientURI *string, contacts []string, frontchannelLogoutSessionRequired *bool, frontchannelLogoutURI *string, grantTypes []string, implicitGrantAccessTokenLifespan *string, implicitGrantIDTokenLifespan *string, jwks map[string]interface{}, jwksURI *string, jwtBearerGrantAccessTokenLifespan *string, logoURI *string, metadata map[string]interface{}, policyURI *string, postLogoutRedirectUris []string, redirectUris []string, responseTypes []string, scope *string, sectorIdentifierURI *string, subjectType *string, tokenEndpointAuthMethod *string, tokenEndpointAuthSigningAlgorithm *string, tosURI *string, userinfoSignedResponseAlgorithm *string, loginBindings *model.LoginBindingsInput) (*model.OAuth2Client, error) {
	return r.C.CreateOAuth2Client(ctx, clients.HydraOperationUpdate, allowedCorsOrigins, audience, authorizationCodeGrantAccessTokenLifespan, authorizationCodeGrantIDTokenLifespan, authorizationCodeGrantRefreshTokenLifespan, backChannelLogoutSessionRequired, backChannelLogoutURI, clientCredentialsGrantAccessTokenLifespan, &clientID, clientName, clientSecret, clientSecretExpiresAt, clientURI, contacts, frontchannelLogoutSessionRequired, frontchannelLogoutURI, grantTypes, implicitGrantAccessTokenLifespan, implicitGrantIDTokenLifespan, jwks, jwksURI, jwtBearerGrantAccessTokenLifespan, logoURI, metadata, policyURI, postLogoutRedirectUris, redirectUris, responseTypes, scope, sectorIdentifierURI, subjectType, tokenEndpointAuthMethod, tokenEndpointAuthSigningAlgorithm, tosURI, userinfoSignedResponseAlgorithm, loginBindings)
}

// DeleteOAuth2Client is the resolver for the deleteOAuth2Client field.
func (r *mutationResolver) DeleteOAuth2Client(ctx context.Context, clientID string) (*model.OAuth2Client, error) {
	return r.C.DeleteOAuth2Client(ctx, clientID)
}

// Owner is the resolver for the owner field.
func (r *oAuth2ClientResolver) Owner(ctx context.Context, obj *model.OAuth2Client) (*string, error) {
	// panic(fmt.Errorf("not implemented: Owner - owner"))
	return nil, nil // TODO: implement
}

// LoginBindings is the resolver for the loginBindings field.
func (r *oAuth2ClientResolver) LoginBindings(ctx context.Context, obj *model.OAuth2Client) (*model.LoginBindings, error) {
	return r.C.GetOAuth2ClientLoginBindings(ctx, *obj.ClientID)
}

// ListOAuth2Clients is the resolver for the listOAuth2Clients field.
func (r *queryResolver) ListOAuth2Clients(ctx context.Context) ([]*model.OAuth2Client, error) {
	return r.C.ListOAuth2Clients(ctx)
}

// GetOAuth2Client is the resolver for the getOAuth2Client field.
func (r *queryResolver) GetOAuth2Client(ctx context.Context, clientID string) (*model.OAuth2Client, error) {
	return r.C.GetOAuth2Client(ctx, clientID)
}

// LoginBindings returns generated.LoginBindingsResolver implementation.
func (r *Resolver) LoginBindings() generated.LoginBindingsResolver { return &loginBindingsResolver{r} }

// OAuth2Client returns generated.OAuth2ClientResolver implementation.
func (r *Resolver) OAuth2Client() generated.OAuth2ClientResolver { return &oAuth2ClientResolver{r} }

type loginBindingsResolver struct{ *Resolver }
type oAuth2ClientResolver struct{ *Resolver }

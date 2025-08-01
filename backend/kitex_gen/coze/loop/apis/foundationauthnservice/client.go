// Code generated by Kitex v0.13.1. DO NOT EDIT.

package foundationauthnservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	authn "github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/foundation/authn"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreatePersonalAccessToken(ctx context.Context, req *authn.CreatePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.CreatePersonalAccessTokenResponse, err error)
	DeletePersonalAccessToken(ctx context.Context, req *authn.DeletePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.DeletePersonalAccessTokenResponse, err error)
	UpdatePersonalAccessToken(ctx context.Context, req *authn.UpdatePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.UpdatePersonalAccessTokenResponse, err error)
	GetPersonalAccessToken(ctx context.Context, req *authn.GetPersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.GetPersonalAccessTokenResponse, err error)
	ListPersonalAccessToken(ctx context.Context, req *authn.ListPersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.ListPersonalAccessTokenResponse, err error)
	VerifyToken(ctx context.Context, req *authn.VerifyTokenRequest, callOptions ...callopt.Option) (r *authn.VerifyTokenResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFoundationAuthNServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFoundationAuthNServiceClient struct {
	*kClient
}

func (p *kFoundationAuthNServiceClient) CreatePersonalAccessToken(ctx context.Context, req *authn.CreatePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.CreatePersonalAccessTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePersonalAccessToken(ctx, req)
}

func (p *kFoundationAuthNServiceClient) DeletePersonalAccessToken(ctx context.Context, req *authn.DeletePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.DeletePersonalAccessTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePersonalAccessToken(ctx, req)
}

func (p *kFoundationAuthNServiceClient) UpdatePersonalAccessToken(ctx context.Context, req *authn.UpdatePersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.UpdatePersonalAccessTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePersonalAccessToken(ctx, req)
}

func (p *kFoundationAuthNServiceClient) GetPersonalAccessToken(ctx context.Context, req *authn.GetPersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.GetPersonalAccessTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPersonalAccessToken(ctx, req)
}

func (p *kFoundationAuthNServiceClient) ListPersonalAccessToken(ctx context.Context, req *authn.ListPersonalAccessTokenRequest, callOptions ...callopt.Option) (r *authn.ListPersonalAccessTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPersonalAccessToken(ctx, req)
}

func (p *kFoundationAuthNServiceClient) VerifyToken(ctx context.Context, req *authn.VerifyTokenRequest, callOptions ...callopt.Option) (r *authn.VerifyTokenResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyToken(ctx, req)
}

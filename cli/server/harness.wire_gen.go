// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject && harness
// +build !wireinject,harness

package server

import (
	"context"

	"github.com/harness/gitness/harness"
	"github.com/harness/gitness/harness/auth/authn"
	"github.com/harness/gitness/harness/auth/authz"
	"github.com/harness/gitness/harness/client"
	"github.com/harness/gitness/harness/router/translator"
	"github.com/harness/gitness/internal/api/controller/repo"
	"github.com/harness/gitness/internal/api/controller/serviceaccount"
	"github.com/harness/gitness/internal/api/controller/space"
	"github.com/harness/gitness/internal/api/controller/user"
	"github.com/harness/gitness/internal/cron"
	"github.com/harness/gitness/internal/router"
	"github.com/harness/gitness/internal/server"
	"github.com/harness/gitness/internal/store/database"
	"github.com/harness/gitness/internal/store/memory"
	"github.com/harness/gitness/types"
)

// Injectors from harness.wire.go:

func initSystem(ctx context.Context, config *types.Config) (*system, error) {
	requestTranslator := translator.ProvideRequestTranslator()
	systemStore := memory.New(config)
	db, err := database.ProvideDatabase(ctx, config)
	if err != nil {
		return nil, err
	}
	userStore := database.ProvideUserStore(db)
	harnessConfig, err := harness.LoadConfig()
	if err != nil {
		return nil, err
	}
	serviceJWTProvider, err := client.ProvideServiceJWTProvider(harnessConfig)
	if err != nil {
		return nil, err
	}
	tokenClient, err := client.ProvideTokenClient(serviceJWTProvider, harnessConfig)
	if err != nil {
		return nil, err
	}
	userClient, err := client.ProvideUserClient(serviceJWTProvider, harnessConfig)
	if err != nil {
		return nil, err
	}
	serviceAccountClient, err := client.ProvideServiceAccountClient(serviceJWTProvider, harnessConfig)
	if err != nil {
		return nil, err
	}
	serviceAccountStore := database.ProvideServiceAccountStore(db)
	authenticator, err := authn.ProvideAuthenticator(userStore, tokenClient, userClient, harnessConfig, serviceAccountClient, serviceAccountStore)
	if err != nil {
		return nil, err
	}
	aclClient, err := client.ProvideACLClient(serviceJWTProvider, harnessConfig)
	if err != nil {
		return nil, err
	}
	authorizer := authz.ProvideAuthorizer(aclClient)
	spaceStore := database.ProvideSpaceStore(db)
	repoStore := database.ProvideRepoStore(db)
	controller := repo.NewController(authorizer, spaceStore, repoStore, serviceAccountStore)
	spaceController := space.NewController(authorizer, spaceStore, repoStore, serviceAccountStore)
	tokenStore := database.ProvideTokenStore(db)
	serviceaccountController := serviceaccount.NewController(authorizer, serviceAccountStore, spaceStore, repoStore, tokenStore)
	userController := user.NewController(authorizer, userStore, tokenStore)
	apiHandler := router.ProvideAPIHandler(systemStore, authenticator, controller, spaceController, serviceaccountController, userController)
	gitHandler := router.ProvideGitHandler(repoStore, authenticator)
	webHandler := router.ProvideWebHandler(systemStore)
	routerRouter := router.ProvideRouter(requestTranslator, apiHandler, gitHandler, webHandler)
	serverServer := server.ProvideServer(config, routerRouter)
	nightly := cron.NewNightly()
	serverSystem := newSystem(serverServer, nightly)
	return serverSystem, nil
}
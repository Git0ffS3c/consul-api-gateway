package e2e

import (
	"context"

	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
)

func SetUpStack(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	var err error
	kindClusterName := envconf.RandomName("consul-api-gateway-test", 30)
	namespace := envconf.RandomName("test", 16)

	for _, f := range []env.Func{
		SetClusterName(kindClusterName),
		SetNamespace(namespace),
		CrossCompileProject,
		BuildDockerImage,
		CreateKindCluster(kindClusterName),
		envfuncs.CreateNamespace(namespace),
		InstallGatewayCRDs,
		CreateServiceAccount(namespace),
		CreateTestConsulContainer(kindClusterName, namespace),
		CreateConsulACLPolicy,
		CreateConsulAuthMethod(namespace),
		InstallConsulAPIGatewayCRDs,
		CreateTestGatewayServer(namespace),
		LoadKindDockerImage(kindClusterName),
	} {
		ctx, err = f(ctx, cfg)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

func TearDownStack(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	var err error
	for _, f := range []env.Func{
		DestroyTestGatewayServer,
		envfuncs.DeleteNamespace(Namespace(ctx)),
		DestroyKindCluster(ClusterName(ctx)),
	} {
		ctx, err = f(ctx, cfg)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}
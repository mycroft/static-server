// A generated module for StaticServer functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
)

type StaticServer struct{}

const (
	GO_VERSION = "1.20"
)

// Use the helloWorld module to show "hello world"
func (s *StaticServer) HelloWorld(ctx context.Context) (string, error) {
	return dag.HelloWorld().ContainerEchoHelloWorld().Stdout(ctx)
}

// Simply build the project with go build
func (m *StaticServer) Build(ctx context.Context, source *Directory) (string, error) {
	return dag.Golang().
		Build(GO_VERSION, source).
		Stdout(ctx)
}

// Build & run tests
func (m *StaticServer) Test(ctx context.Context, source *Directory) (string, error) {
	return dag.Golang().
		Test(GO_VERSION, source).
		Stdout(ctx)
}

// Run the whole pipeline: test, build & publish
func (m *StaticServer) Publish(ctx context.Context, regUsername string, regPassword *Secret, regAddress string, imageName string, source *Directory, tag string) (string, error) {
	_, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}

	address, err := dag.Golang().Build(GO_VERSION, source).
		WithRegistryAuth(regAddress, regUsername, regPassword).
		Publish(ctx, fmt.Sprintf("%s/%s:%s", regAddress, imageName, tag))
	if err != nil {
		return "", err
	}

	return address, err
}

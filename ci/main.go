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

// Returns a container for testing & build
func (m *StaticServer) getContainer(source *Directory) *Container {
	return dag.Container().
		From(fmt.Sprintf("golang:%s-alpine", GO_VERSION)).
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"apk", "add", "curl"})
}

// Build static-server
func (m *StaticServer) build(source *Directory) *Container {
	return m.getContainer(source).
		WithExec([]string{"go", "build"})
}

// Simpl ybuild the project with go build
func (m *StaticServer) Build(ctx context.Context, source *Directory) (string, error) {
	return m.build(source).
		Stdout(ctx)
}

// Run tests
func (m *StaticServer) Test(ctx context.Context, source *Directory) (string, error) {
	return m.getContainer(source).
		WithExec([]string{"go", "test"}).
		Stdout(ctx)
}

// Build & test the project
func (m *StaticServer) BuildTest(ctx context.Context, source *Directory) (string, error) {
	_, err := m.Build(ctx, source)
	if err != nil {
		return "", err
	}

	return m.Test(ctx, source)
}

// Run the whole pipeline: test, build & publish
func (m *StaticServer) Publish(ctx context.Context, regUsername string, regPassword *Secret, regAddress string, imageName string, source *Directory, tag string) (string, error) {
	_, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}

	address, err := m.build(source).
		WithRegistryAuth(regAddress, regUsername, regPassword).
		Publish(ctx, fmt.Sprintf("%s/%s:%s", regAddress, imageName, tag))
	if err != nil {
		return "", err
	}

	return address, err
}

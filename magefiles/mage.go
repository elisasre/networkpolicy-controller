//go:build mage

package main

import (
	"context"
	"os"

	"github.com/elisasre/mageutil"
)

const (
	AppName   = "networkpolicy-controller"
	RepoURL   = "https://github.com/elisasre/networkpolicy-controller"
	ImageName = "quay.io/elisaoyj/networkpolicy-controller"
)

// Build binaries for executables under ./cmd
func Build(ctx context.Context) error {
	return mageutil.BuildAll(ctx)
}

// UnitTest whole repo
func UnitTest(ctx context.Context) error {
	return mageutil.UnitTest(ctx)
}

// IntegrationTest whole repo
func IntegrationTest(ctx context.Context) error {
	return mageutil.IntegrationTest(ctx, "./cmd/"+AppName)
}

// Lint all go files.
func Lint(ctx context.Context) error {
	return mageutil.LintAll(ctx)
}

// VulnCheck all go files.
func VulnCheck(ctx context.Context) error {
	return mageutil.VulnCheckAll(ctx)
}

// LicenseCheck all files.
func LicenseCheck(ctx context.Context) error {
	return mageutil.LicenseCheck(ctx, os.Stdout, mageutil.CmdDir+AppName)
}

// Clean removes all files ignored by git
func Clean(ctx context.Context) error {
	return mageutil.Clean(ctx)
}

// Image creates docker image.
func Image(ctx context.Context) error {
	return mageutil.DockerBuildDefault(ctx, ImageName, RepoURL)
}

// PushImage creates docker image.
func PushImage(ctx context.Context) error {
	return mageutil.DockerPushAllTags(ctx, ImageName)
}

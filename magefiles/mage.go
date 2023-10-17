//go:build mage

package main

import (
	"context"
	"os"

	"github.com/elisasre/mageutil"
	"github.com/magefile/mage/mg"
)

const (
	AppName   = "networkpolicy-controller"
	RepoURL   = "https://github.com/elisasre/networkpolicy-controller"
	ImageName = "quay.io/elisaoyj/networkpolicy-controller"
)

type (
	Go        mg.Namespace
	Docker    mg.Namespace
	Workspace mg.Namespace
)

// Build binaries for executables under ./cmd
func (Go) Build(ctx context.Context) error {
	return mageutil.BuildAll(ctx)
}

// UnitTest runs unit tests for whole repo
func (Go) UnitTest(ctx context.Context) error {
	return mageutil.UnitTest(ctx)
}

// Lint runs lint for all go files
func (Go) Lint(ctx context.Context) error {
	return mageutil.LintAll(ctx)
}

// VulnCheck runs vuln check for all packages
func (Go) VulnCheck(ctx context.Context) error {
	return mageutil.VulnCheckAll(ctx)
}

// LicenseCheck checks licences of all packages
func (Go) LicenseCheck(ctx context.Context) error {
	return mageutil.LicenseCheck(ctx, os.Stdout, mageutil.CmdDir+AppName)
}

// Clean removes all files ignored by git
func (Workspace) Clean(ctx context.Context) error {
	return mageutil.Clean(ctx)
}

// Image creates docker image
func (Docker) Image(ctx context.Context) error {
	return mageutil.DockerBuildDefault(ctx, ImageName, RepoURL)
}

// PushImage pushes docker image
func (Docker) PushImage(ctx context.Context) error {
	return mageutil.DockerPushAllTags(ctx, ImageName)
}

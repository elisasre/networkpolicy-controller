//go:build mage

package main

import (
	"context"
	"os"

	goutil "github.com/elisasre/mageutil/golang"
	"github.com/magefile/mage/mg"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
	//mage:import
	_ "github.com/elisasre/mageutil/golangcilint/target"
	//mage:import
	_ "github.com/elisasre/mageutil/govulncheck/target"
	//mage:import
	_ "github.com/elisasre/mageutil/golicenses/target"
	//mage:import
	docker "github.com/elisasre/mageutil/docker/target"
	//mage:import
	golang "github.com/elisasre/mageutil/golang/target"
)

// Configure imported targets
func init() {
	os.Setenv(mg.VerboseEnv, "1")
	os.Setenv("CGO_ENABLED", "0")

	golang.BuildTarget = "./cmd/networkpolicy-controller"
	docker.ImageName = "quay.io/elisaoyj/networkpolicy-controller"
	docker.ProjectUrl = "https://github.com/elisasre/networkpolicy-controller"
}

type Go mg.Namespace

// UnitCoverProfile create coverage profile in text format
//
// TODO: Once integration tests are added remove this and use importedgo target.
func (Go) UnitCoverProfile(ctx context.Context) error {
	return goutil.CreateCoverProfile(ctx, goutil.CombinedCoverProfile, goutil.UnitTestCoverDir)
}

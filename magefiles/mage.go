//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
	//mage:import go
	_ "github.com/elisasre/mageutil/tool/golangcilint"
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
	docker.ImageName = "europe-north1-docker.pkg.dev/sose-sre-5737/sre-public/networkpolicy-controller"
	docker.ProjectUrl = "https://github.com/elisasre/networkpolicy-controller"
}

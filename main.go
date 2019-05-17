package main

import (
	"github.com/cv21/gen-generator-mock/generator"

	"github.com/cv21/gen/pkg"
	plugin "github.com/hashicorp/go-plugin"
)

const (
	pluginRepoURL = "github.com/cv21/gen-generator-mock@v1.0.0"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		Plugins: map[string]plugin.Plugin{
			pluginRepoURL: &pkg.NetRPCWorker{Impl: generator.NewMockGenerator()},
		},
	})
}

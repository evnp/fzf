package main

import (
	"github.com/evnp/fzfx/src"
	"github.com/evnp/fzfx/src/protector"
)

var version string = "0.27"
var revision string = "devel"

func main() {
	protector.Protect()
	fzfx.Run(fzfx.ParseOptions(), version, revision)
}

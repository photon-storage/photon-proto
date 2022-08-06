//go:build tools
// +build tools

package pbt

// Trick go mod into requiring protoc-gen-go-cast and therefore Gazelle won't prune it.
import (
	_ "github.com/photon-storage/protoc-gen-go-cast"
)

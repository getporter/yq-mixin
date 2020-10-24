//go:generate packr2
package yq

import (
	"get.porter.sh/porter/pkg/context"
)

const defaultClientVersion string = "3.4.1"

type Mixin struct {
	*context.Context
	ClientVersion string
}

// New azure mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		Context:       context.New(),
		ClientVersion: defaultClientVersion,
	}
}

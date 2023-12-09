package rpc

import (
	"github.com/google/wire"
)

// ProviderSet is rpc providers.
var ProviderSet = wire.NewSet(NewClient)

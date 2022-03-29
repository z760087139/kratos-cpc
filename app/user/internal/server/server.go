package server

import (
	"github.com/google/wire"
)

// ProviderSet is policy providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)

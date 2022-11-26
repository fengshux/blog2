package db

import "github.com/google/wire"

// ProviderSet is controller providers.
var ProviderSet = wire.NewSet(
	NewPG,
)

package types

const (
	// ModuleName defines the module name
	ModuleName = "coinactions"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_coinactions"
)

var (
	ParamsKey = []byte("p_coinactions")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmoschain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmoschain"

    
)

var (
	ParamsKey = []byte("p_cosmoschain")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "mymodule"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mymodule"

    
)

var (
	ParamsKey = []byte("p_mymodule")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}

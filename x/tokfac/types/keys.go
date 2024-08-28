package types

const (
	// ModuleName defines the module name
	ModuleName = "tokfac"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokfac"
)

var (
	ParamsKey = []byte("p_tokfac")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

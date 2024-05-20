package types

const (
	// ModuleName defines the module name
	ModuleName = "hssn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hssn"
)

var (
	ParamsKey = []byte("p_hssn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

var (
	TokenCountKey = []byte{0x01}
	TokenKey      = []byte{0x02}
)

func GetTokenKey(id uint64) []byte {
	return append(TokenKey, UInt64Bytes(id)...)
}

package types

var (
	LaunchpadCountKey        = []byte{0x01}
	LaunchpadKey             = []byte{0x02}
	LaunchpadDepositCountKey = []byte{0x03}
	LaunchpadDepositKey      = []byte{0x04}
)

func GetLaunchpadKey(id uint64) []byte {
	return append(LaunchpadCountKey, UInt64Bytes(id)...)
}

func GetDepositToLaunchpadKey(id, count uint64) []byte {
	return append(LaunchpadDepositKey, append(UInt64Bytes(count), UInt64Bytes(id)...)...)
}

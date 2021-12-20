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

func GetDepositToLaunchpadKey(id uint64) []byte {
	return append(LaunchpadDepositKey, UInt64Bytes(id)...)
}

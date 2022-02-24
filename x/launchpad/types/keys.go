package types

var (
	LaunchpadCountKey        = []byte{0x01}
	LaunchpadKey             = []byte{0x02}
	LaunchpadCreatedKey      = []byte{0x03}
	LaunchpadActiveKey       = []byte{0x04}
	LaunchpadSuccessKey      = []byte{0x05}
	LaunchpadFailKey         = []byte{0x06}
	LaunchpadDepositCountKey = []byte{0x07}
	LaunchpadDepositKey      = []byte{0x08}
)

func GetLaunchpadKey(id uint64) []byte {
	return append(LaunchpadKey, UInt64Bytes(id)...)
}

func GetDepositToLaunchpadKey(id, count uint64) []byte {
	return append(LaunchpadDepositKey, append(UInt64Bytes(count), UInt64Bytes(id)...)...)
}

func GetLaunchpadCreatedKey() []byte {
	return append(LaunchpadCreatedKey, []byte("CREATED")...)
}

func GetLaunchpadActiveKey() []byte {
	return append(LaunchpadActiveKey, []byte("ACTIVE")...)
}

func GetLaunchpadSuccessKey() []byte {
	return append(LaunchpadSuccessKey, []byte("SUCCESS")...)
}

func GetLaunchpadFailKey() []byte {
	return append(LaunchpadFailKey, []byte("FAIL")...)
}

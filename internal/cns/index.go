package cns

// Platforms
const (
	PlatformUndefined = 0
	PlatformAndroid   = 1
	PlatformIOS       = 2
	PlatformWeb       = 3
)

func PlatformIsValid(v int) bool {
	return v == PlatformUndefined ||
		v == PlatformAndroid ||
		v == PlatformIOS ||
		v == PlatformWeb
}

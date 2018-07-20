package constants

// Constants for database key.
const (
	KeySourcePrefix = "so:"

	KeyServerPrefix = "s:"
)

// FormatSourceKey will format a source key.
func FormatSourceKey(source string) []byte {
	return []byte(KeySourcePrefix + source)
}

// FormatServerKey will format a server key.
func FormatServerKey(source, server string) []byte {
	return []byte(KeyServerPrefix + source + ":" + server)
}

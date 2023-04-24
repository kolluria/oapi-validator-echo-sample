package utils

func TypeToPtr[T any](v T) *T {
	return &v
}

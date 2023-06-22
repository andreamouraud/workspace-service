package utils

func Map[T, U any](arr []T, f func(T) U) []U {
	res := make([]U, len(arr))
	for i := range arr {
		res[i] = f(arr[i])
	}
	return res
}

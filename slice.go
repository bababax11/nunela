package nunela

func every[T any](s []T, f func(int, T) bool) bool {
	for i, v := range s {
		if !f(i, v) {
			return false
		}
	}
	return true
}

func mapSlice[T any, U any](sl []T, f func(T) U) []U {
	out := make([]U, 0, len(sl))
	for _, v := range sl {
		out = append(out, f(v))
	}
	return out
}

func rangeSlice(len int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = i
	}
	return s
}

func keys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for key := range m {
		s = append(s, key)
	}
	return s
}

// func equal[T comparable](xs []T, ys []T) bool {
// 	if len(xs) != len(ys) {
// 		return false
// 	}
// 	for i := range xs {
// 		if xs[i] != ys[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

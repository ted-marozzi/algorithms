package leetcode

func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

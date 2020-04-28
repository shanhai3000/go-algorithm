package sort

type Comparable interface {
	CompareTo(comparable Comparable) int
}


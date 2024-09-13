package set

type Set interface {
	Capacity() int
	Size() int
	Has(key any) bool
	Add(key any)
	Remove(key any)
	Iterate() <-chan any
	Clear()
}

type Map interface {
	Set
	Get(key any) any
	AddKV(key, value any)
	IterateValues() <-chan any
}

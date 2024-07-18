package collection

type Map[K int | string, V any] interface {

	// Get element.
	Get(K) V

	Put(K, V) Map[K, V]

	Contains(K) bool
}

func NewHashMap[K int | string, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		Values: make(map[K]V),
	}
}

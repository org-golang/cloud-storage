package collection

type HashMap[K int | string, V any] struct {
	Values map[K]V
}

func (h *HashMap[K, V]) Get(key K) V {
	return h.Values[key]
}

func (h *HashMap[K, V]) Put(key K, value V) Map[K, V] {
	h.Values[key] = value

	return h
}

func (h *HashMap[K, V]) Contains(key K) bool {

	_, ok := h.Values[key]

	return ok
}

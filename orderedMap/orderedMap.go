package orderedMap

type Orderable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		uintptr | string
}

type OrderedMap[K Orderable, V any] struct {
	keys   []K
	values []V
}

func NewOrderMap[K Orderable, V any]() (ret OrderedMap[K, V]) {
	ret.Init()
	return ret
}

func (om *OrderedMap[K, V]) Init() {
	om.keys = make([]K, 0)
	om.values = make([]V, 0)
}

func (om *OrderedMap[K, V]) GetKeys() []K {
	return om.keys
}

func (om *OrderedMap[K, V]) GetValues() []V {
	return om.values
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.keys)
}

func (om *OrderedMap[K, V]) GetIndex(key K) (int, bool) {
	right := om.Len()
	if right == 0 {
		return 0, false
	}
	left := 0
	index := right / 2
	for left < right {
		pKey := om.keys[index]
		if pKey == key {
			return index, true
		} else if pKey < key {
			left = index + 1
		} else if pKey > key {
			right = index
		}
		index = (left + right) / 2
	}
	return index, false
}

func (om *OrderedMap[K, V]) GetWithOK(k K) (ret V, ok bool) {
	index, ok := om.GetIndex(k)
	if ok {
		ret = om.values[index]
	}
	return ret, ok
}

func (om *OrderedMap[K, V]) Get(k K) (ret V) {
	index, ok := om.GetIndex(k)
	if ok {
		ret = om.values[index]
	}
	return ret
}

func (om *OrderedMap[K, V]) GetByIndex(index int) (k K, v V, ok bool) {
	if index < 0 || index >= om.Len() {
		return k, v, false
	}
	return om.keys[index], om.values[index], true
}

func (om *OrderedMap[K, V]) GetKeyByIndex(index int) (k K, ok bool) {
	if index < 0 || index >= om.Len() {
		return k, false
	}
	return om.keys[index], true
}

func (om *OrderedMap[K, V]) GetValueByIndex(index int) (v V, ok bool) {
	if index < 0 || index >= om.Len() {
		return v, false
	}
	return om.values[index], true
}

func (om *OrderedMap[K, V]) insertAt(k K, v V, index int) {
	om.keys = append(append(om.keys[:index], k), om.keys[index:]...)
	om.values = append(append(om.values[:index], v), om.values[index:]...)
}

func (om *OrderedMap[K, V]) Insert(k K, v V) bool {
	index, ok := om.GetIndex(k)
	if ok {
		return false
	}
	om.insertAt(k, v, index)
	return true
}

func (om *OrderedMap[K, V]) Set(k K, v V) {
	index, ok := om.GetIndex(k)
	if ok {
		om.values[index] = v
	} else {
		om.insertAt(k, v, index)
	}
}

func (om *OrderedMap[K, V]) Delete(k K) bool {
	index, ok := om.GetIndex(k)
	if !ok {
		return false
	}
	om.keys = append(om.keys[:index], om.keys[index+1:]...)
	om.values = append(om.values[:index], om.values[index+1:]...)
	return true
}

func (om *OrderedMap[K, V]) DeleteByIndex(index int) bool {
	if index < 0 || index >= om.Len() {
		return false
	}
	om.keys = append(om.keys[:index], om.keys[index+1:]...)
	om.values = append(om.values[:index], om.values[index+1:]...)
	return true
}

func (om *OrderedMap[K, V]) Range(f func(K, V)) {
	for index, k := range om.keys {
		f(k, om.values[index])
	}
}

func (om *OrderedMap[K, V]) RangeKeys(f func(K)) {
	for _, k := range om.keys {
		f(k)
	}
}

func (om *OrderedMap[K, V]) RangeValues(f func(V)) {
	for _, v := range om.values {
		f(v)
	}
}

package taskfile

type DeepCopier[T any] interface {
	DeepCopy() T
}

func deepCopySlice[T any](orig []T) []T {
	if orig == nil {
		return nil
	}
	c := make([]T, len(orig))
	for i, v := range orig {
		if copyable, ok := any(v).(DeepCopier[T]); ok {
			c[i] = copyable.DeepCopy()
		} else {
			c[i] = v
		}
	}
	return c
}

func deepCopyMap[K comparable, V any](orig map[K]V) map[K]V {
	if orig == nil {
		return nil
	}
	c := make(map[K]V, len(orig))
	for k, v := range orig {
		if copyable, ok := any(v).(DeepCopier[V]); ok {
			c[k] = copyable.DeepCopy()
		} else {
			c[k] = v
		}
	}
	return c
}

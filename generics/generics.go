package generics

// Delete 删除对应下标索引的元素
func Delete[T any](idx int, slice []T) []T {
	if idx >= len(slice) || idx < 0 {
		panic("idx out of range slice")
	}
	var newSlice []T
	if len(slice)-1 < (cap(slice)+1)/2 {
		newSlice = make([]T, len(slice)-1, (cap(slice)+1)/2)
		for i, j := 0, 0; i < len(slice); i++ {
			if i == idx {
				continue
			}
			newSlice[j] = slice[i]
			j++
		}
	} else {
		newSlice = slice[:idx]
		for i := idx + 1; i < len(slice); i++ {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

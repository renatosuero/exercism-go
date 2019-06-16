package listops

//IntList type defines an array of integers
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

//Foldr method folds the list (right)
func (i IntList) Foldr(f binFunc, value int) int {
	for ii := len(i) - 1; ii >= 0; ii-- {
		value = f(i[ii], value)
	}
	return value
}

//Foldl method folds the list (left)
func (i IntList) Foldl(f binFunc, value int) int {
	for _, ii := range i {
		value = f(value, ii)
	}
	return value
}

//Filter method applies function based in the filter
func (i IntList) Filter(f predFunc) IntList {
	list := make(IntList, 0, len(i))
	for _, j := range i {
		even := f(j)
		if even {
			list = append(list, j)
		}
	}
	return list
}

//Length method returns the length of the list
func (i IntList) Length() (len int) {
	for range i {
		len++
	}
	return len
}

//Map method applies function in all elements
func (i IntList) Map(f unaryFunc) IntList {
	list := make(IntList, 0, len(i))
	for _, j := range i {
		list = append(list, f(j))
	}
	return list
}

//Reverse method returns list inverted
func (i IntList) Reverse() IntList {
	list := make(IntList, 0, len(i))
	for j := len(i) - 1; j >= 0; j-- {
		list = append(list, i[j])
	}
	return list
}

//Append method returns new list with new values added
func (i IntList) Append(value IntList) IntList {
	for _, j := range value {
		i = append(i, j)
	}
	return i
}

//Concat method returns a new list with new values concatenated
func (i IntList) Concat(value []IntList) IntList {
	list := make(IntList, 0, len(i))
	list = append(list, i...)
	for _, j := range value {
		list = append(list, j...)
	}
	return list
}

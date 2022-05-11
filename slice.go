package oper

type Number interface {
	~int8 | ~int16 | ~int | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type List interface {
	Number | ~string
}

func In[N List](fat []N, sub N) bool {
	_, ok := InI(fat, sub)
	return ok
}

func Index[N List](fat []N, idx N) int {
	index, _ := InI(fat, idx)
	return index
}

func Extend[N List](fat []N, sub []N) []N {

	fat = append(fat, sub...)

	return fat
}

func Pop[N List](fat []N, idx int) []N {
	if idx >= len(fat) || len(fat)+idx < 0 {
		return nil
	}

	if idx < 0 {
		idx = len(fat) + idx
	}

	return append(fat[:idx], fat[(idx+1):]...)
}

func Remove[N List](fat []N, value N) []N {
	for i, v := range fat {
		if v == value {
			fat = append(fat[:i], fat[(i+1):]...)
		}
	}

	return fat
}

func Append[N List](fat []N, value N) []N {

	return append(fat, value)
}

func Insert[N List](fat []N, idx int, value N) []N {
	var res = make([]N, len(fat), len(fat)+1)
	copy(res, fat)

	res = append(res[:idx], value)

	return append(res, fat[idx:]...)
}

func Count[N List](fat []N, value N) (count int) {

	for _, v := range fat {
		if v == value {
			count += 1
		}
	}

	return count
}

func InI[N List](fat []N, sub N) (int, bool) {

	for i, v := range fat {
		if v == sub {
			return i, true
		}
	}

	return -1, false
}

func Equal[N List](fat []N, sub []N) bool {
	if len(fat) != len(sub) {
		return false
	}
	for i, n := range fat {
		if n != sub[i] {
			return false
		}
	}
	return true
}

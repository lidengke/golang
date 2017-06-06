package test

// SliceTest just for slice example test
func SliceTest() {
	var a []int
	appSlice(a, 1)

}

func appSlice(a []int, elment int) []int {
	return append(a, elment)
}

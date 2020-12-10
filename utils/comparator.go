package utils

// Comparator func
type Comparator func(interface{}, interface{}) int

func intAscComparator(x, y interface{}) int {
	return x.(int) - y.(int)
}

func intDescComparator(x, y interface{}) int {
	return y.(int) - x.(int)
}

func stringAscComparator(x, y interface{}) int {
	if x == y {
		return 0
	} else if x.(string) < y.(string) {
		return -1
	} else {
		return 1
	}
}

func stringDescComparator(x, y interface{}) int {
	if x == y {
		return 0
	} else if x.(string) < y.(string) {
		return 1
	} else {
		return -1
	}
}

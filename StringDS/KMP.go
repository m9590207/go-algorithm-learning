package StringDS

func nextTable(p string) []int {
	sizeP := len(p)
	nextArr := make([]int, sizeP)
	nextArr[0] = -1
	i, j := 0, -1

	for i < sizeP-1 {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			nextArr[i] = j
		} else {
			j = nextArr[j]
		}
	}
	return nextArr
}

func Kmp(src, p string) int {
	if p == "" {
		return 0
	}

	nextArr := nextTable(p)
	i, j := 0, 0
	sizeSrc, sizeP := len(src), len(p)

	for i < sizeSrc && j < sizeP {
		if j == -1 || src[i] == p[j] {
			i++
			j++
		} else {
			j = nextArr[j]
		}
	}

	if j == sizeP {
		return i - sizeP
	} else {
		return -1
	}
}

package StringDS

import "testing"

func TestKMP(t *testing.T) {
	src, p := "abaacababcac", "ababc"
	if Kmp(src, p) != 5 {
		t.Errorf("src = %s  p = %s 回傳不為5", src, p)
	}

	src, p = "abc", ""
	if Kmp(src, p) != 0 {
		t.Errorf("src = %s  p = %s 回傳不為0", src, p)
	}

	src, p = "aaaaa", "bba"
	if Kmp(src, p) != -1 {
		t.Errorf("src = %s  p = %s 回傳不為-1", src, p)
	}
}

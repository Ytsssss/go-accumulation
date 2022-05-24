package algorithm

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int, 0)
	for _, char := range s {
		if i, ok := m[char]; ok {
			m[char] = i + 1
		} else {
			m[char] = 1
		}
	}
	for _, char := range t {
		if i, ok := m[char]; ok && i > 0 {
			m[char] = i - 1
			if i-1 == 0 {
				delete(m, char)
			}
		} else {
			return false
		}
	}

	return len(m) == 0
}

func isAnagram2(s string, t string) bool {
	nums1, nums2 := make([]int, 26), make([]int, 26)
	for _, i := range s {
		nums1[i-'a']++
	}
	for _, i := range t {
		nums2[i-'a']++
	}
	for i := 0; i < 26; i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

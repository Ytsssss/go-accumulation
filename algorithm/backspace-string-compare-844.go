package algorithm

func BackspaceCompare(s string, t string) bool {
	i, j, counti, countj := len(s)-1, len(t)-1, 0, 0
	for i >= 0 && j >= 0 {
		if s[i] == '#' {
			counti++
			i--
			continue
		}
		if t[j] == '#' {
			countj++
			j--
			continue
		}
		if counti > 0 {
			i--
			counti--
			continue
		}
		if countj > 0 {
			j--
			countj--
			continue
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	for i >= 0 {
		if s[i] == '#' {
			counti++
			i--
		} else if counti > 0 {
			counti--
			i--
		} else {
			return false
		}
	}
	for j >= 0 {
		if t[j] == '#' {
			countj++
			j--
		} else if countj > 0 {
			countj--
			j--
		} else {
			return false
		}
	}
	return true
}

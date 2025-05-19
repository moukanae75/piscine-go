package piscine

func Map(f func(int) bool, a []int) []bool {
	table := make([]bool, len(a))

	for i, c := range a {
		if f(c) {
			table[i] = true
		} else {
			table[i] = false
		}
	}
	return table
}

package main

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{-2, 3, 4, 7, 11, 15}
	c := intersection(a, b)

	for _, val := range c {
		println(val)
	}
}

// amazed this is not a part of standard library
func contains(roles []int, role int) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}

func intersection(a []int, b []int) []int {

	c := []int{}

	for _, val := range a {
		if contains(b, val) {
			c = append(c, val)
		}
	}

	return c
}


//Реализовать пересечение двух неупорядоченных множеств.
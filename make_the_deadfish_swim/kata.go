package make_the_deadfish_swim

// https://www.codewars.com/kata/51e0007c1f9378fa810002a9/train/go
func Parse(data string) []int {

	v := 0
	output := make([]int, 0)

	for i := range data {
		switch string(data[i]) {
		case "i":
			v++
			break
		case "s":
			v = v * v
			break
		case "d":
			v--
			break
		case "o":
			output = append(output,v)
			break
		}
	}

	return output
}

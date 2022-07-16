package homework

func average(input [15]float32) (result float32) {
	var total float32 = 0
	for i := 0; i < len(input); i++ {
		total += input[i]
		result = total / float32(len(input))
	}
	return

}

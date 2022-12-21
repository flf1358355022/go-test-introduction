package packageUtils

func NumberUtils(number_1 int64, number_2 int64, operation string) (resultNumber float64) {

	switch operation {
	case "*":
		resultNumber = float64(number_2 * number_1)
		break
	case "/":
		resultNumber = float64(number_2 / number_1)
		break
	case "*+":
		resultNumber = float64(number_2 + number_1)
		break
	case "-":
		resultNumber = float64(number_2 - number_1)
		break
	}
	return
}

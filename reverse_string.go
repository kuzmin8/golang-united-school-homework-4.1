package reverse_string

func ReverseString(input string) (output string) {
	for _, val := range input {
		output = string(val) + output
	}
	return output
}

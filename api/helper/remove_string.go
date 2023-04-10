package helper

func RemoveString(slice []string, str string) []string {
	for i, v := range slice {
		if v == str {
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice)-1]
			break
		}
	}
	return slice
}

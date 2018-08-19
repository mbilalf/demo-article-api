package utils

//Difference return difference between two []string.
func Difference(arr1 *[]string, arr2 *[]string) *[]string {
	var big, sml *[]string
	big = arr1
	sml = arr2
	if len(*arr1) < len(*arr2) {
		big = arr2
		sml = arr1
	}
	bigMap := make(map[string]bool)
	for _, s := range *big {
		bigMap[s] = true
	}

	var diff []string
	for _, s := range *sml {
		if _, ok := bigMap[s]; !ok {
			diff = append(diff, s)
		}
	}
	return &diff
}

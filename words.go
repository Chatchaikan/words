package words

import "strings"
import "fmt"

func WordCount(s string) map[string]int {
	fmt.Println("Teed!")

	strs := strings.Fields(s)
	res := make(map[string]int)

	for _, str := range strs {
		res[strings.ToLower(str)]++
	}
	return res
	//return map[string]int{"x": 1}
}

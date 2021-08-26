package main

import (
	"fmt"
	"sort"
	"strings"
)

func seachKw(data string) string {
	word := strings.Split(data, "")
	sort.Strings(word)
	return strings.Join(word, "")
}

func main() {

	data := make(map[string][]string)
	result := [][]string{}
	listAnagram := []string{"dani", "nida", "rusak", "kasur", "tusuk", "kusut"}

	for _, v := range listAnagram {
		index := seachKw(v)
		data[index] = append(data[index], v)

	}

	//each group
	for _, j := range data {
		result = append(result, j)
	}
	fmt.Println(result)

}

package prettier

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func Standardise(s string) string {
	return strings.Title(strings.ToLower(s))
}

func MakeStringArrayCommaSeparatedString(stringArr []string) string {
	str := ""
	for _, s := range stringArr {
		if str != "" {
			str += ", "
		}
		str += s
	}
	return str
}

func PrintKeyValueArray(a []KeyValue) {
	for _, elem := range a {
		fmt.Println(elem.Key + ": " + strconv.Itoa(elem.Value))
	}
}

func Print2DArray(arr [][]string) {
	for _, record := range arr {
		for i, elem := range record {
			fmt.Print(elem)
			if i < len(record)-1 {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
}

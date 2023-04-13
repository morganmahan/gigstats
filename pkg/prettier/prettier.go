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

func PrintKeyValueArray(a []KeyValue) {
	for _, elem := range a {
		fmt.Println(elem.Key + ": " + strconv.Itoa(elem.Value))
	}
}

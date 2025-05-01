package prettier

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/morganmahan/gigstats/internal/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type KeyValue struct {
	Key   string
	Value int
}

func Standardise(s string) string {
	caser := cases.Title(language.English)
	return caser.String(strings.ToLower(s))
}

func StandardiseCommaSeparated(str string) string {
	strArr := strings.Split(str, ", ")
	for i, s := range strArr {
		strArr[i] = Standardise(s)
	}
	return strings.Join(strArr, ", ")
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

func PrintGigsArray(gigs []types.Gig) {
	for i, gig := range gigs {
		fmt.Printf("%d. %s - %s - %s - %s\n", i+1, MakeStringArrayCommaSeparatedString(gig.Bands), gig.Venue, gig.Date, MakeStringArrayCommaSeparatedString(gig.Who))
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

func FlattenArray(valsToSplit [][]string) []string {
	vals := []string{}
	for _, value := range valsToSplit {
		vals = append(vals, value...)
	}
	return vals
}

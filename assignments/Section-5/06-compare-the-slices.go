package section5

import (
	"fmt"
	"sort"
	"strings"
)

func compareTheSlices() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	namesC := strings.Split(namesA, ", ")

	if len(namesB) == len(namesC) {
		sort.Strings(namesB)
		sort.Strings(namesC)
		for index := range namesC {
			if namesC[index] != namesB[index] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}

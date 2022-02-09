package section5

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	header = "Location,Size,Beds,Baths,Price"
	data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
	separator = ","
)

var (
	locations                  []string
	sizes, beds, baths, prices []int
)

func housingPricesAverages() {
	for _, row := range strings.Split(data, "\n") {
		cols := strings.Split(row, separator)
		locations = append(locations, cols[0])

		si, err := strconv.Atoi(cols[1])
		sizes = append(sizes, si)

		be, err := strconv.Atoi(cols[2])
		beds = append(beds, be)

		ba, err := strconv.Atoi(cols[3])
		baths = append(baths, ba)

		pr, err := strconv.Atoi(cols[4])
		prices = append(prices, pr)

		if err != nil {
			fmt.Printf("Error while string to Int conversion \n Error message %s", err)
		}
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range strings.Split(data, "\n") {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}
	fmt.Printf("%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s", "")

	var sum int

	for _, n := range sizes {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(sizes)))

	sum = 0
	for _, n := range beds {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(beds)))

	sum = 0
	for _, n := range baths {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(baths)))

	sum = 0
	for _, n := range prices {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(prices)))
	fmt.Printf("\n%s", strings.Repeat("=", 75))
}

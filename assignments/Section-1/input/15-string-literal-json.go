package stringLiteralJSON

import "fmt"

func stringLiteralJSON() {
	json :=
		`
		{
			"Items": [{
				"Item": {
					"name": "Teddy Bear"
				}
			}]
		}
	`
	fmt.Println(json)
}

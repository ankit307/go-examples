package main

import (
	"fmt"
)

func main() {
	days := make(map[int]string)
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	fmt.Printf("%v\n", days)

	//Deletes key
	delete(days, 1)
	fmt.Printf("%v\n", days)

	fmt.Printf("%v\n", len(days))

}

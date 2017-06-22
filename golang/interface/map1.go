package main

import (
	"fmt"
)

func main() {
	rating := map[string]float32{"C": 5, "go": 4.5, "python": 4.5, "C++": 2, "C#": 1}
	charpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is", charpRating)
	} else {
		fmt.Println("we have no rating associcated with C# in the map")
	}
	delete(rating, "C")
	fmt.Println(rating)
}

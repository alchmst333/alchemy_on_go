/*
Problem:

Create a Go program that initializes a slice of strings representing a list of fruits. Your task is to iterate over the slice and print each fruit's name along with its position in the list.

Details:

Initialize a slice of strings with at least 5 different fruit names.
Use a for loop with range to iterate over the slice.
For each fruit in the slice, print the position (index) and the fruit's name in the format: Position: Fruit Name.
Example Output:

If the slice contains the fruits ["Apple", "Banana", "Cherry", "Date", "Elderberry"], the output should be:

Position: 0, Fruit: Apple
Position: 1, Fruit: Banana
Position: 2, Fruit: Cherry
Position: 3, Fruit: Date
Position: 4, Fruit: Elderberry
*/

package main

import "fmt"

var fruit = []string{"papaya", "mango", "cherry", "lychee", "durian"}

func main() {
	for i, v := range fruit {
		fmt.Printf("Position:%d, Fruit:%s\n", i, v)
	}
}

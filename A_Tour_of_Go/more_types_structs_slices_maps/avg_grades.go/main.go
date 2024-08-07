/*
Exercise: Student Grades
Create a program that manages a list of students and their grades. The program should:
1. Define a Student struct with fields for the student's name and a slice of their grades.
2. Create a slice of Student instances.
3. Iterate over the slice using a range loop to print each student's name and their average grade.

Detailed Steps:
1. Define a Student struct with Name and Grades fields.
2. Create a slice of Student instances with at least 3 students, each having a few grades.
3. Write a function to calculate the average grade of a student.
4. Use a range loop to iterate over the slice of students and print each student's name along with their average grade.
*/

package main

import "fmt"

type Students struct {
	Name   string
	Grades []int
}

// Write a func to calculate the average grades of a student
func avg(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func main() {
	students := []Students{
		{"Mars", []int{89, 72, 98, 85}},
		{"Gaia", []int{92, 68, 91, 83}},
		{"Atlas", []int{63, 74, 87, 100}},
	}

	// Iterate over the slice using a range loop
	for _, student := range students {
		avgGrade := avg(student.Grades)
		fmt.Printf("Student: %s, Average Grade: %.2f\n", student.Name, avgGrade)
	}
}

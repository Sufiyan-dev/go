package main

import "fmt"

func GenDisplaceFn(a, vi, si float64) func(t float64) float64 {
	return func(t float64) float64 {
		s := (0.5*a*t*t + vi*t + si)
		return s
	}
}

func main() {

	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement")
	fmt.Scan(&acceleration, &initialVelocity, &initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println("Enter time ")
	fmt.Scan(&time)

	fmt.Println("Displacement for time", time, " is : ", fn(time))

}

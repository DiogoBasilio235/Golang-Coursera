package main

import (
	"fmt"
	"math"
	"time"
)

func GenDisplaceFunc(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
	return fn
}

func Delay(delay float64) {
	sec := int(delay)
	time.Sleep(time.Duration(sec) * time.Second)
}

func main() {

	var in_acceleration float64
	var in_init_velocity float64
	var in_init_displacement float64
	var in_time float64

	fmt.Println("Please enter the accelatarion:")
	fmt.Scan(&in_acceleration)

	fmt.Println("Please enter the initial velocity:")
	fmt.Scan(&in_init_velocity)

	fmt.Println("Please enter the initial displacement:")
	fmt.Scan(&in_init_displacement)

	fmt.Println("Please enter the time:")
	fmt.Scan(&in_time)

	Delay(in_time)

	fmt.Println("The displacement is:", GenDisplaceFunc(in_acceleration, in_init_velocity, in_init_displacement)(in_time))

}

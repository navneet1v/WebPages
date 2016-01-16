package main

import( 
	"fmt"
	//"math"
)
/*
	if new var use := to assign them a new value
	else use  =  to assign them value
*/

/*
	x int
	p *int
	a [3]int
*/
// we can return two values in this language

func swap(a, b int) (int,int) {
	return b,a
}

// function to add two numbers
func addNumbers(a, b int) int {
	return a + b 
}


func main(){	
	//fmt.Println("This is just a ")
	//fmt.Printf("The square root of %d is %g",7,math.Sqrt(7));
	
	x := 17
	y := 23

	fmt.Println(addNumbers(x, y))
	x, y = swap(17, 23)
	x = 12
	fmt.Println(x, y)

	var a , b , c int
	a = 10
	b = 12
	c = 23
	fmt.Println(a,b,c)
}
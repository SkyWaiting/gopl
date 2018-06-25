//https://blog.csdn.net/u012855229/article/details/51766314
package tempconv

import "fmt"

func Example_one() {
	{
		//!+arith
		fmt.Printf("%g\n", BoilingC-FreezingC)
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC))
		//!-arith
	}

	/*
		//!+arith
		fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch
		//!-arith
	*/

	// Output:
	// 100
	// 180
}

func Example_two() {
	//!+printf
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	//!-printf

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
	// 100
	// 100
}

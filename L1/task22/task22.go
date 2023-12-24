package task22

import (
	"fmt"
	"math/big"
)

func Launch() {
	a := new(big.Int)
	a.SetString("12000000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(a.Text(10))
	b := new(big.Int)
	b.SetString("11000000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(b.Text(10))

	//The result is
	c := new(big.Int)
	c.Add(a, b)
	fmt.Printf("a+b: %s\n", c.Text(10))
	c.Mul(a, b)
	fmt.Printf("a*b: %s\n", c.Text(10))
	c.Div(a, b)
	fmt.Printf("a/b: %s\n", c.Text(10))

}

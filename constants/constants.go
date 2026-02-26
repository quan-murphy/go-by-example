
/*
const delcares a constant value 

numeric constant has no type until given one, such as by explicit conversion 

number can be given a type by using it in a context that requires one, suc has varialbe assignment or function call

*/
package main 

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n 
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}



/*

for is GoLang's only looping construct 


Classic initial/condition/after `for` loop

Another way of achieving basic "do this N times" iteration is range over an integer 

for without condition will run forever without `break` or `return`

Can use `continue` to next interation of the loop
*/

package main 

import("fmt")


func main(){
	i :=1 
	for i < 3{
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++{
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}



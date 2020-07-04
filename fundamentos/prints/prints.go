package main

import "fmt"

func main() {
	fmt.Print("Opa!")
	fmt.Print(" linha!")

	fmt.Println(" Nova!")
	fmt.Println("linha!")

	x := 3.141516

	// fmt.Println("O valor de x e " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x e " + xs)
	fmt.Println("O valor de x e ", x)

	fmt.Printf("O valor de x e %2.f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.v %v %v", a, b, b, c, d)

}

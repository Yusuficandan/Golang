package main

func main() {

	// 1: int x, float 64 y type conversion sample

	/*
		x := 75
		var y float64
		y = float64(x)

		fmt.Println(y)
	*/

	// 2: multiple assing sample x,y =y,x

	/*
		x := 5
		y := 10

		fmt.Println("X:", x, "Y:", y)

		x, y = y, x

		fmt.Println("X:", x, "Y:", y)

	*/

	// 3: non english variable names
	/*

		yaş := 40
		fmt.Println(yaş)

	*/

	// 4: shadowing kavramı

	/*
		x := 5

			if true {
				x := 10
				fmt.Println(x)
			}

			fmt.Println(x)
	*/

	// 5: 40 as a string
	/*
		x := 40
		s := string(x)

		fmt.Printf("%v, %T\n", x, x)

		fmt.Printf("%v, %T\n", s, s)

		y := strconv.Itoa(x)

		fmt.Printf("%v, %T\n", y, y)
	*/

}

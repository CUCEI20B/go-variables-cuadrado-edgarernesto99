package main

import "fmt"

func main()  {
	var lado uint;
	fmt.Scan(&lado);

	area := lado*lado;
	fmt.Println(area);
}
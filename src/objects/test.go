package objects

import "fmt"

func Test(){
	var point1 = Point{10, 20}
	fmt.Println(point1);

	var point2 = &point1;
	point2.X = 11
	point2.Y = 12
	fmt.Println(point1)
	fmt.Println(point2)

	point3 := GetInstance()
	fmt.Println(point3)
	point3.SetX(3);
	point3.SetY(4);
	fmt.Println(point3)

	point4 := GetInstance()
	fmt.Println(point4)
}


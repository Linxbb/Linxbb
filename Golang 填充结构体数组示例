package main

import "fmt"

type RoomOccupancys struct {
	RoomOccupancyList []RoomOccupancy
}

type RoomOccupancy struct {
	RoomCount  int
	AdultCount int 
	ChildCount int
	CotsCount  int
}

func main() {

	test1 := RoomOccupancys{[]RoomOccupancy{{5, 4, 3,5}, {4, 4, 4, 4}}}
	Test(test1)

}
func Test(data RoomOccupancys) {
	fmt.Println(data.RoomOccupancyList)
	var a,b,c int
	for _,one := range data.RoomOccupancyList{
		a += one.RoomCount
		b += one.AdultCount
		c += one.ChildCount
		fmt.Printf("%+v\n",one)
	}
	fmt.Println(a,b,c)


}

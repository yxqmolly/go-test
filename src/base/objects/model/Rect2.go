package model


type Point struct{
	X,Y int
}
type Rect2 struct{
	Leftup, RightDown *Point
}
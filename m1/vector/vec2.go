package vector

type Vec2 struct {
	X int
	Y int
}

func (self *Vec2) Negate() {
	self.X = -self.X
	self.Y = -self.Y
}

func MakeVec2(x,y int) *Vec2 {
	return &Vec2{x, y}
}

func AddVec2(v1 *Vec2, v2 *Vec2) *Vec2 {
	return &Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

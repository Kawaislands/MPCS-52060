package vector

type Vec2 struct {
	X int // x-coordinate
	Y int // y-coordinate
}

// AddVec2 adds two vectors together and returns a new Vector
func AddVec2(v1 Vec2, v2 Vec2) Vec2 {
	return Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

func (self *Vec2) Negate() {
	self.X = -self.X
	self.Y = -self.Y
}

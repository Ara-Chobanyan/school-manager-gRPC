package app

// A lot of data structure in this project are going be the same; My guess in a actually production project you would, make these data structure as libraries
// Student data structure
type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Grade        float32
	Behavior     float32
	Satisfaction float32
	Critical     bool
	Exceptional  bool
	Tutelary     bool
  Staff bool
}

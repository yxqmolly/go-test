package model

/**
 * 结构体的定义,首字母小写
 */
type student struct {
	id int
	name string
	score float64
}

func (s *student) getScore() float64 {
	return s.score
}
func (s *student) setScore(a float64) {
	s.score = a
}
func NewStudent2(id int, n string, s float64) *student {
	return &student{
		id: id,
		name : n,
		score: s,
	}
}


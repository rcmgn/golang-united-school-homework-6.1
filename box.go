package golang_united_school_homework

import "errors"

const idxOutOfRange = "index out of the box capacity range"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("shapes quantity out of the box capacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity {
		return nil, errors.New(idxOutOfRange)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (s Shape, e error) {
	if i >= b.shapesCapacity {
		return nil, errors.New(idxOutOfRange)
	}
	s = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= b.shapesCapacity {
		return nil, errors.New(idxOutOfRange)
	}
	res := b.shapes[i]
	b.shapes[i] = shape
	return res, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (result float64) {
	for _, sh := range b.shapes {
		result += sh.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (result float64) {
	for _, sh := range b.shapes {
		result += sh.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() (e error) {
	e = errors.New("circle is not exist in the list")
	var res []Shape
	for _, sh := range b.shapes {
		switch sh.(type) {
		case Circle:
			e = nil
		default:
			res = append(res, sh)
		}
	}
	b.shapes = res
	return e
}

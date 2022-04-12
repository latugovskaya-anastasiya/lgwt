package smi_test

import (
	"testing"

	smi "github.com/latugovskaya-anastasiya/lgwt/struct_methods_interfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := smi.Rectangle{Width: 10.0, Height: 10.0}
	got := smi.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   smi.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: smi.Rectangle{Width: 10, Height: 10}, hasArea: 100.0},
		{name: "Circle", shape: smi.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: smi.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}

		})
	}
}

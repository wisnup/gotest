package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("expected %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%#v got %g want %g", shape, want, got)
		}
	}

	// table driven test
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triange", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, test := range areaTest {
		t.Run(test.name, func(t *testing.T) {

			checkArea(t, test.shape, test.hasArea)
		})
	}
}

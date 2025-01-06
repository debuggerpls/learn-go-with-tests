package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(&Rectangle{10, 10})
	want := 40.0

	assert(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		shape := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, shape, want)
	})
	t.Run("circles", func(t *testing.T) {
		shape := Circle{10}
		want := math.Pi * 10 * 10

		checkArea(t, shape, want)
	})
}

func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12.0, 6.0}, want: 72.0},
		{shape: Circle{10}, want: math.Pi * 10 * 10},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

func assert(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

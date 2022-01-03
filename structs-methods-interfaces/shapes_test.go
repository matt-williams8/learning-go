package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

var RECTANGLE = Rectangle{Height: 2.4, Width: 5.3}

func floatsMatch(tb testing.TB, expected float64, got float64) {
	tb.Helper()
	if got != expected {
		tb.Errorf("expected %v, got %v", expected, got)
	}
}

func TestPerimeter(t *testing.T) {
	got := RECTANGLE.Perimeter()
	expected := 15.4

	floatsMatch(t, expected, got)
}

func TestArea(t *testing.T) {

	checkArea := func(tb testing.TB, shape Shape, expectedArea float64) {
		tb.Helper()
		floatsMatch(tb, expectedArea, shape.Area())
	}

	areaTests := []struct {
		testName     string
		shape        Shape
		expectedArea float64
	}{
		{testName: "Rectangle", shape: RECTANGLE, expectedArea: 12.72},
		{testName: "Circle", shape: Circle{4}, expectedArea: 50.27},
		{testName: "Triangle", shape: Triangle{Base: 10, Height: 5}, expectedArea: 25},
	}

	for _, areaTest := range areaTests {
		t.Run(fmt.Sprintf("Should calculate area for %v", areaTest.testName), func(t *testing.T) {
			checkArea(t, areaTest.shape, areaTest.expectedArea)
		})
	}
}

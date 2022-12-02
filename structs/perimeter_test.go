package perimiter

import "testing"

func TestPerimeter(t *testing.T) {
	ret := Retangle{10.0, 10.0}
	got := Perimeter(ret)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		ret := Retangle{12.0, 6.0}
// 		want := 72.0

// 		checkArea(t, ret, want)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793

// 		checkArea(t, circle, want)
// 	})
// }

func TestArea(t *testing.T) {
	testCases := []struct {
		desc     string
		shape    Shape
		expected float64
	}{
		{
			desc:     "rectangles",
			shape:    Retangle{12.0, 6.0},
			expected: 72.0,
		},
		{
			desc:     "circles",
			shape:    Circle{10},
			expected: 314.1592653589793,
		},
		{
			desc:     "triangle",
			shape:    Triangle{12, 6},
			expected: 36.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.shape.Area()

			if got != tc.expected {
				t.Errorf("%#v got %g want %g", tc.shape, got, tc.expected)
			}
		})
	}
}

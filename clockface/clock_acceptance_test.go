package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/mrayone/learn-go/clockface"
)

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{
			simpleTime(0, 0, 0),
			clockface.Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l clockface.Line, ls []clockface.Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

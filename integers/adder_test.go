package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' bug got '%d'", expected, sum)
	}
}

// Example of code to generate docs
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}

package concurrencypatterns_test

import "github.com/mrayone/learn-go/concurrencypatterns"

func ExampleGenerator() {
	concurrencypatterns.Generator()
	//Output:
	//"Counter at :  0"
	//"Counter at :  1"
	//"Counter at :  2"
	//"Counter at :  3"
	//"Counter at :  4"
	//Done with Counter
}

func ExampleGeneratorPosition() {
	concurrencypatterns.GeneratorPosition()
	//Output:
	//
}

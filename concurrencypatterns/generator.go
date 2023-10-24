package concurrencypatterns

import "fmt"

func foo() <-chan string {
	myChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			myChannel <- fmt.Sprintf("%s %d", "Counter at : ", i)
		}
	}()

	return myChannel
}

func Generator() {
	myChannel := foo()

	for i := 0; i < 5; i++ {
		fmt.Printf("%q\n", <-myChannel)
	}

	fmt.Println("Done with Counter")
}

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
		}
	}()

	return positionChannel
}

func GeneratorPosition() {
	pc1 := updatePosition("Legolas :")
	pc2 := updatePosition("Gandalf :")

	for i := 0; i < 5; i++ {
		fmt.Println(<-pc1) // blocked until pc2 receives
		fmt.Println(<-pc2) // blocked until pc1 receives
	}

	fmt.Println("Done with getting update on positions.")
}

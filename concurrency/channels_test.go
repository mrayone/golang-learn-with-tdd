package concurrency_test

import "github.com/mrayone/learn-go/concurrency"

func ExampleRun() {
	concurrency.Run()
	//Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleRunCloseAndChannel() {
	concurrency.RunCloseAndChannel()
	//Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 0
}

func ExampleRunCloseAndChannelBreak() {
	concurrency.RunCloseAndChannelBreak()
	//Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleRunAlongRange() {
	concurrency.RunAlongRange()
	//Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

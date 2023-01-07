package main

import (
	"os"
	"time"

	"github.com/mrayone/learn-go/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

package api

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, personToGreet string) {
	fmt.Fprintf(writer, "Hello, %s!", personToGreet)
}

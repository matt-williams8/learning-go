// REFRESH PACKAGE -> MODULE KNOWLEDGE, HOW CAN I HAVE FILES IN DIFFERENT PACKAGES IN THE SAME MODULE
package main

import (
	"dependency-injection/api"
	"os"
)

func main() {
	api.Greet(os.Stdout, "Test name")
}

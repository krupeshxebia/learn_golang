package main

import (
	"fmt"
	"math"
	"github.com/krupeshxebia/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Pow(3, 2))

	fmt.Println(strutil.Reverse("Hello World"))
}

package main

import "flag"
import "fmt"
import "strconv"
import "os"

var UsageCalled bool

func main() {
	UsageCalled = false
	/* you can use the following and set the default to false
	 * when you do this and then just pass -b
	 * into the program it gets set to true

	 */
	var myBool bool
	flag.BoolVar(&myBool, "b", false, "-b=true|-b=false")
	var fileNameArg string
	flag.StringVar(&fileNameArg, "f", "unk", "-f=filename")
	flag.Parse()
	fmt.Println("the value of b")
	fmt.Println(strconv.FormatBool(myBool))
	fmt.Println(fileNameArg)
	//flag.PrintDefaults()

	// if we pass in unk paramters then these don't get run
	// so an implicit exit()
	if UsageCalled {
		fmt.Println("UsageCalled")
	}
	fmt.Println("Do we get here if bad args")
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	UsageCalled = true
}

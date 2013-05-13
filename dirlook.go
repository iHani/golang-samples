/**
 * just print out the contents and some metadata about directory contents
 **/

package main

import "log"
import "fmt"
import iou "io/ioutil"


 func main() {
 	dir :=  "/Volumes/HD2/temp"
	fmt.Println(dir +"\n====================")
	// dirArray []os.FileInfo
	// error err
	dirInfo, err := iou.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(len(dirInfo))
	fmt.Println(" items in the folder " + dir)
	fmt.Println("=========================================")
	for i := range dirInfo {
		println(dirInfo[i].IsDir(), dirInfo[i].Name() + " ")
	}


}
/**
 * just print out the contents and some metadata about directory contents
 *
 * tabWriter does not yet work correctly
 **/

package main

import "log"
import "fmt"
import iou "io/ioutil"
import "text/tabwriter"
import "os"
 import "strings"

//import "time"

 func main() {
 	dir :=  "/Volumes/HD2/temp"
 	w := new(tabwriter.Writer);
	w.Init(os.Stdout, 31, 2, 2, ':', tabwriter.AlignRight)


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

	for  i := range dirInfo {
		var isDir string
		if dirInfo[i].IsDir() {
			isDir = "D" 
		} else {
			isDir = "-"
		}
		var myStr = strings.Join([]string{
			isDir,
			//string(dirInfo[i].Size()),
			//string(dirInfo[i].Mode()),
			//dirInfo[i].ModTime().String(),
			dirInfo[i].Name()}, 
			"\t")

		fmt.Fprintln(w, myStr)	
	}
	w.Flush()


}
package main

import "io/ioutil"
import "fmt"
import "log"
import "os"
import "strconv"

/** some simple file operations
 *  note that for larger files you would want to use bufffered input and output
 **/
func main() {

	fileName := "./temp/tempfile.txt"
	writeToFile(fileName)
	readFromFile(fileName)
	pwd()
	writeManyStringToFile("./temp/lgfile.txt", 10)
	prtFileStat(fileName)

}

//================================================================
func readFromFile(fileName string) {
	msg("reading from small file: " + fileName)
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(contents))
	tmp := make([]byte, len(contents))
	for i := 0; i < len(contents); i++ {
		tmp[i] = contents[i]
	}
	fmt.Println(string(tmp) + ":end of contents")
}

//================================================================
func prtFileStat(fileName string) {
	msg("getting file stat info for: " + fileName)
	fi, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("   Name: ", fi.Name())
	fmt.Println("   Size: ", strconv.FormatInt(fi.Size(), 10))
	fmt.Println("   Mode: ", fi.Mode().String())
	fmt.Println("ModTime: ", fi.ModTime().String())
	fmt.Println("  IsDir: ", strconv.FormatBool(fi.IsDir()))
	fmt.Println("Sys: ", "Not printing out Sys() because it may not be printable")
	fmt.Println("==========")
}

/**
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
**/

//================================================================
func msg(msg string) {

	fmt.Println("==========\n", msg, "\n==========")
}

//================================================================
func writeToFile(fileName string) {
	// write data to file

	fileString := "this is a line of text\n"
	msg("Writing to file: " + fileName)

	/* if file does not exist it is created (but folders have to exist)
	 * if called multiple times this overwrites the file (dones not append)
	 */
	err := ioutil.WriteFile(fileName, []byte(fileString), 0644)
	if err != nil {
		log.Println("Error writing to file")
		log.Fatal(err)
	}
}

//================================================================
func pwd() {
	//curent working dir
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	msg("Current working dir: " + cwd)
}

//================================================================
func writeManyStringToFile(fileName string, i int) {

	msg("writeManyStringToFile")

	str := "sdfffasklgsgwefsgksrgdrg458938609463njsvdbt^T*&%%^zdfsgdsfsdfsdfsdfsfsfsdfZ\n"
	fo, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	writeBuf := make([]byte, 1024)
	writeBuf = []byte(str)
	n := len(str)
	for j := 0; j < i; j++ {
		if _, err := fo.Write(writeBuf[:n]); err != nil {
			panic(err)
		}
	}

}

//================================================================

//================================================================

//================================================================

//================================================================

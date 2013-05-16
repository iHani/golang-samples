package main

//TODO consider pkg path/filepath/glob

/** this i useless and needs to be replaced with a
 * dir walker
 **/

import "log"
import "fmt"
import "os"
import iou "io/ioutil"

func main() {
	f := new(FileAndFolder)
	err := f.init("/Volumes/HD2/temp", true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.ParentFolder)
	fmt.Println(f.FileArray)
	fmt.Println(f.FolderArray)
}

type FileAndFolder struct {
	ParentFolder   string
	FileArray      []os.FileInfo
	FolderArray    []os.FileInfo
	RecurseFolders bool
}

// need the * in *FileAndFolder
func (f *FileAndFolder) init(s string, r bool) error {
	f.RecurseFolders = r
	dirInfo, err := iou.ReadDir(s)
	if err != nil {
		//TODO how to handle the errors here
		return err
	}
	f.ParentFolder = s
	for i := range dirInfo {
		if dirInfo[i].IsDir() {
			f.FolderArray = append(f.FolderArray, dirInfo[i])
		} else {
			f.FileArray = append(f.FileArray, dirInfo[i])
		}
	}
	return nil
}

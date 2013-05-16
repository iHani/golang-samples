package main

//http://weekly.golang.org/doc/go1.html#path_filepath
import "fmt"
import "path/filepath"
import "os"
import "log"

/**
* okay this works really well
*
* you can recurse the folder if you want or not
* just set Recurse = true|false
*
**/

//TODO define the walk function outside of main body and return it

func main() {
	OrgPath := "/Volumes/HD2/temp"
	Recurse := true
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() && path != OrgPath && !Recurse {
			fmt.Println("skipping dir:", path)
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	}
	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {
		log.Fatal(err)
	}

}

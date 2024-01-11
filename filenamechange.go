package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type file struct {
	path string
	name string
	dir  string
}


// Really basic recursive rename, I just added a
// word to the file name instead of parsing and changing
func nameChange() {

	pathArg := "sample"
	fsys := os.DirFS(pathArg)

	var files []file

	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err.Error())
		}

		if d.IsDir() {
			fmt.Println("Skipping dir: ", d)
		} else {
			files = append(files, file{
				path: filepath.Join(pathArg, p),
				name: d.Name(),
				dir:  path.Dir(p),
			})
		}

		return nil
	})
	for _, file := range files {
		newFile := buildNewPath(file, pathArg)
		err := os.Rename(file.path, newFile.path)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func buildNewPath(f file, root string) file {
	if f.dir == "." {
		return file{
			name: f.name,
			path: filepath.Join(root, f.name),
			dir:  f.dir,
		}
	} else {
		return file{
			name: f.name,
			path: filepath.Join(root, f.dir, f.name),
			dir:  f.dir,
		}
	}

}

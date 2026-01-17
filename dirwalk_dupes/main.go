// package main walks a dir and returns a 2D array of dupes
/*

./root/
├── a.txt
├── b.txt
└── foo/
    ├── bar/
    │   └── e.txt
    ├── c.txt
    └── d.txt

a.txt = "xyz"
b.txt = "abc"
c.txt = "xyz"
d.txt = "123"
e.txt = "abc"

Given the above dir structure, return a 2D array of the files with duplicate contents, like below:

[["abc", "root/b.txt", "root/foo/bar/e.txt"], ["xyz", "root/a.txt", "root/foo/c.txt"], ["123", "root/foo/d.txt"]]

*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateStructure generates the specific directory tree and files requested.
func CreateStructure(basePath string) error {
	files := map[string]string{
		"a.txt":         "xyz",
		"b.txt":         "abc",
		"foo/c.txt":     "xyz",
		"foo/d.txt":     "123",
		"foo/bar/e.txt": "abc",
	}

	for path, content := range files {
		fullPath := filepath.Join(basePath, path)

		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}
	}

	return nil
}

func main() {
	root := "root"
	err := CreateStructure(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

package size

import (
	"fmt"
	"gosize/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func ListSizes(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		fmt.Printf("%s (%s)\n", root, utils.FormatSize(info.Size()))
		return nil
	}

	fmt.Printf("%s/\n", root)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil // skip root itself
		}

		relativePath, _ := filepath.Rel(root, path)
		depth := strings.Count(relativePath, string(os.PathSeparator))

		prefix := strings.Repeat("│   ", depth)

		if info.IsDir() {
			fmt.Printf("%s├── %s/\n", prefix, info.Name())
		} else {
			fmt.Printf("%s├── %s (%s)\n", prefix, info.Name(), utils.FormatSize(info.Size()))
		}

		return nil
	})

	return err
}

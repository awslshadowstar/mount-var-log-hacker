package logutil

import (
	"os"
)

func AttachToRoot() {
	err := os.Symlink("/", "/var/log/host/root_link")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Symbolic link created successfully.")
}

func DetachFromRoot() {
	// 符号链接路径
	linkPath := "/var/log/host/root_link"

	// 删除符号链接
	err := os.Remove(linkPath)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Symbolic link deleted successfully.")
}

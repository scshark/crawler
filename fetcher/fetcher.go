package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetcher(Url string) ([]byte, error) {
	time.Sleep(200 * time.Millisecond)
	resp, err := http.Get(Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status is not ok the code is %v", resp.StatusCode)
	}
	body := bufio.NewReader(resp.Body)
	d := determineEncoding(body)
	utf8Reader := transform.NewReader(body, d.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("new reader peek error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


package crawler

import (
	"bufio"
	"fmt"
	"github.com/fatrbaby/imooc-crawler/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string, toUTF8 bool) ([]byte, error) {
	<-rateLimiter

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code: %d\n", response.StatusCode)
	}

	var body io.Reader

	if toUTF8 {
		buff := bufio.NewReader(response.Body)
		body = transform.NewReader(buff, EncodingGuesser(buff).NewDecoder())
	} else {
		body = response.Body
	}

	raw, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	return raw, err
}

func EncodingGuesser(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)

	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

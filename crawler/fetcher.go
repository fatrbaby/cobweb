package crawler

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string, toUTF8 bool) ([]byte, error) {
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
		body = transform.NewReader(response.Body, EncodingGuesser(response.Body).NewDecoder())
	} else {
		body = response.Body
	}

	raw, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	return raw, err
}

func EncodingGuesser(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)

	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

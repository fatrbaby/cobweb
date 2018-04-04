package crawler

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
)

func FetchContentFrom(url string, toUtf8 bool) []byte {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Errorf("Http error code: %d\n", response.StatusCode))
	}

	var body io.Reader

	if toUtf8 {
		body = transform.NewReader(response.Body, EncodingGuesser(response.Body).NewDecoder())
	} else {
		body = response.Body
	}

	raw, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err)
	}



	return raw
}

func EncodingGuesser(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)

	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

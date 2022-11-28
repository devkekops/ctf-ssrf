package client

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/k3a/html2text"
)

type Client interface {
	GetInfo(url string) (string, error)
}

type cli struct {
	httpClient *http.Client
}

func NewCli() Client {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	client := &http.Client{
		Transport: t,
		Timeout:   time.Duration(5 * int(time.Second)),
	}
	return &cli{
		httpClient: client,
	}
}

func extractText(s string) string {
	return html2text.HTML2Text(s)
}

func (c *cli) GetInfo(url string) (string, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		plain := extractText(string(body))

		return plain, nil
	}
	return "", nil
}

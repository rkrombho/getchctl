package getch

import (
	"io/ioutil"
	"net/url"
)

func (c *Client) Encrypt(value string) string {
	//build the encryption URL
	mergedurl := mergeUrl(c.url, "encrypt")
	//POST the value
	resp, err := c.httpClient.PostForm(mergedurl, url.Values{"value": {value}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//read the response
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(cont)
}

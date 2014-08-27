package getch

import (
	"io/ioutil"
	"log"
	"os"
)

func (c *Client) Get(key string) string {
	//Generate the URL to query the given key
	url := mergeUrl(c.url, key)
	//GET it
	resp, err := c.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.Status == "404 Not Found" {
		log.Fatalf("No value found for key %v", key)
		os.Exit(404)
	}
	//read the response
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(cont)
}

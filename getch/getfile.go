package getch

import (
	"io/ioutil"
	"log"
	"os"
)

func (c *Client) GetFile(name, outputname string) {
	//the Getch URL is juat the bas url with the filename appended
	url := mergeUrl(c.url, name)
	//GET the file
	resp, err := c.httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.Status == "404 Not Found" {
		log.Fatalf("No file found on Getch server named %v for the context of this host", name)
		os.Exit(404)
	}
	if resp.Status == "500 Internal Server Error" {
		log.Fatalf("Getch server error while trying to fetch file with name %v. Most of the time this comes down to a syntax error in a template that you try to download. Please check the server side logs.", name)
		os.Exit(500)
	}
	//read it into memory
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//now write the file to the current directory with the given gilename
	var filename string
	if len(outputname) > 0 {
		filename = outputname
	} else {
		filename = name
	}
	//write the file. The permission bitmask 0666 will cuse the 
	//systems default umask to be used
	err = ioutil.WriteFile(filename, cont, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func (r *WebsiteUrl) Read(p []byte) (n int, err error) {
	res, err := http.Get(r.URL)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	doc, err := html.Parse(res.Body)

	if err != nil {
		return 0, err
	}

	data := extractData(doc)

	copy(p, []byte(data))
	return len(data), nil
}

func WebsiteReader(URL string) *WebsiteUrl {
	return &WebsiteUrl{URL: URL}
}

func extractData(n *html.Node) string {
	var data string

	if n.Type == html.TextNode {
		data += n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data += extractData(c)
	}

	return data
}

func main() {
	reader := WebsiteReader("URL IN HERE")

	data := make([]byte, 1024)

	if len(data) < 99999 {
		data = append(data, make([]byte, 99999-len(data))...)
	}

	n, err := reader.Read(data)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	err = os.WriteFile("result.txt", data, 14754)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result Created!")

	fmt.Printf("Read %d\n", n)
}

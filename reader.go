package main

import (
	"fmt"
	"net/http"
)

func (r *WebsiteUrl) Read(p []byte) (n int, err error) {
	res, err := http.Get(r.URL)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	return res.Body.Read(p)
}

func WebsiteReader(URL string) *WebsiteUrl {
	return &WebsiteUrl{URL: URL}
}

func main() {
	reader := WebsiteReader("URL IN HERE!")

	data := make([]byte, 1024)

	n, err := reader.Read(data)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}

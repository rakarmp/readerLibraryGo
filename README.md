## Reader Library Go

Web Reader To Get Data Website

```
Write Pure Go
Go Version 1.21.0
```

### Use

Change Url In reader.go file

```
func main() {
	reader := WebsiteReader("URL IN HERE")

	data := make([]byte, 1024)

	n, err := reader.Read(data)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}
```

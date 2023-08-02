package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cixtor/readability"
)

const html = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>HTML 5 Boilerplate</title>
    <link rel="stylesheet" href="style.css">
  </head>
  <body>
%s
  </body>
</html>
`

func main() {
	filePath := flag.String("file", "", "file path of the html file")
	flag.Parse()

	if *filePath == "" {
		log.Fatalln("file path is required")
		return
	}

	f, err := os.OpenFile(*filePath, os.O_RDONLY, 0o_600)
	if err != nil {
		log.Fatalf("unable to open the file: %v\n", err)
		return
	}

	art, err := loadFile(f)
	if err != nil {
		log.Fatalf("unable to parse the file: %v\n", err)
		return
	}

	fmt.Printf(html, art.Content)
}

func loadFile(r io.Reader) (readability.Article, error) {
	return readability.New().Parse(r, "/")
}

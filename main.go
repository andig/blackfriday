package main

import "github.com/russross/blackfriday"

func main() {
	_ = blackfriday.MarkdownBasic([]byte("foo"))
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var filePath string

func write(message string) {
	fmt.Println("Write " + filePath)
	err := ioutil.WriteFile(filePath, []byte(message), 0644)
	if err != nil {
		panic(err)
	}
}

func append(message string) {
	fmt.Println("Append " + filePath)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, message)
}

func read() []byte {
	fmt.Println("Read " + filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return content
}

func delete() {
	fmt.Println("Delete " + filePath)
	if err := os.Remove(filePath); err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("[Usage]  %s %s \n", filepath.Base(os.Args[0]), "/path/to/file")
		os.Exit(0)
	}

	filePath = os.Args[1]

	// 作成
	write("Text file is created.\n")
	fmt.Println(string(read()))

	// 追記
	append("Message is appended.\n")
	fmt.Println(string(read()))

	// 上書き
	write("Text file is updated.\n")
	fmt.Println(string(read()))

	// 削除
	delete()
}

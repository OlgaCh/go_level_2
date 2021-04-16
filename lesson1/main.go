package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ErrorWithTimestamp struct {
	text      string
	timestamp string
}

func New(text string) error {
	return &ErrorWithTimestamp{
		text:      text,
		timestamp: time.Now().String(),
	}
}

func (e *ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("error: %s\noccur at: %s \n", e.text, e.timestamp)
}

func handlePanic() {
	defer func() {
		if v := recover(); v != nil {
			err := New("Error from panic")
			fmt.Printf("Handle panic with details\n%v\n", err)
		}
	}()

	fmt.Println("### A panic example ###")
	var arr []int
	fmt.Println(5 / len(arr))
}

func createAndCloseFile(index int) {
	f, _ := os.Create(fmt.Sprintf("files/File_%d.txt", index))
	defer f.Close()
}

func createEmptyFiles() {
	fmt.Println("### Creating 1M of empty files ###")

	n := 1_000_000
	dir := "files"

	_ = os.Mkdir(dir, 0700)
	for i := 0; i < n; i++ {
		createAndCloseFile(i + 1)
	}
	files, _ := ioutil.ReadDir(dir)
	fmt.Printf("Created %v files.\n", len(files))

	fmt.Println("Cleaning up...\n")
	os.RemoveAll(dir)
}

func panicInGoroutine() {
	fmt.Println("### Handling panic in goroutine ###")

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("Recovered from panic in goroutine:", v)
			}
		}()

		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)

}

func main() {
	handlePanic()
	createEmptyFiles()
	panicInGoroutine()
}

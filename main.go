package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func dnld(epizode int, path string) error {
	pathRemote := "https://golangshow.com/cdn/episodes/"
	if epizode < 10 {
		pathRemote += string("00") + fmt.Sprint(epizode) + ".mp3"
	}
	if epizode >= 10 && epizode < 100 {
		pathRemote += string("0") + fmt.Sprint(epizode) + ".mp3"
	}
	if epizode >= 100 && epizode < 1000 {
		pathRemote += fmt.Sprint(epizode) + ".mp3"
	}
	if epizode < 0 || epizode >= 1000 {
		log.Fatal("No valid epizod number")
	}

	resp, err := http.Get(pathRemote)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	} else {
		fmt.Println("downloaded from URL", pathRemote)
	}
	return nil
}

func main() {
	for i := 1; i < 100; i++ {
		localPath := "files/" + fmt.Sprint(i) + ".mp3"
		err := dnld(i, localPath)
		if err != nil {
			fmt.Println(err)
		}
	}

}

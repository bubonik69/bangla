package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func dnld(epizode int, path string, wg *sync.WaitGroup) {
	pathRemote := "https://golangshow.com/cdn/episodes/"
	pathRemote += fmt.Sprintf("%03d", epizode) + ".mp3"
	defer wg.Done()
	if epizode < 0 || epizode > 123 {
		log.Fatal("No valid epizod number")
	}

	resp, err := http.Get(pathRemote)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("downloaded from URL", pathRemote)
	}
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		localPath := "files/" + fmt.Sprint(i) + ".mp3"
		wg.Add(1)
		go dnld(i, localPath, &wg)
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("time is:", duration)
}

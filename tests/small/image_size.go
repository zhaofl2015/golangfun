package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://cdn.17zuoye.com/fs-resource/585395dc2ed9b636fe24ed91.png")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	g := m.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	// The resolution is height x width
	resolution := height * width
	fmt.Println(height, width)

	// Print results
	fmt.Println(resolution, "pixels")
}

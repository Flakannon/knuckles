package ascii

import (
	"log"
	"os"
)

func GetKnucklesArt() string {
	file, err := os.Open("internal/ascii/knuckles.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, 2048)
	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	return string(buf[:n])
}

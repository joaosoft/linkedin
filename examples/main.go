package main

import (
	"github.com/joaosoft/linkedin"
)

func main() {
	_, err := linkedin.NewLinkedin()
	if err != nil {
		panic(err)
	}
}

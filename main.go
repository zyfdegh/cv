package main

import (
	"flag"
	"log"
)

func main() {
	conf := flag.String("f", "local.json", "cv config file")
	flag.Parse()

	cmd, err := NewCVGenCmd(*conf)
	if err != nil {
		log.Fatalln(err)
	}
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start

type ShowNesting struct {
	Str string `xml:"a>b>str"`
}

// :show end

func printXML(v interface{}) {
	d, err := xml.Marshal(v)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("XML: %s\n\n", string(d))
}

func main() {
	// :show start
	v := &ShowNesting{
		Str: "str",
	}
	printXML(v)

	// :show end
}

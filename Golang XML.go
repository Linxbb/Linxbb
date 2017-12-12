package main

import (
	"github.com/stackerzzq/xj2go"
)

func main() {
	xmlParse := xj2go.New("./response.xml", "Contents", "")
	xmlParse.XMLToGo()
}

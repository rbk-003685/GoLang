package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
  myuuid := uuid.NewV4()
  fmt.Println(myuuid)
}

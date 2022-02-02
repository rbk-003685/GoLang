package main

import (
  "fmt"

  "github.com/google/uuid"
)

func main(){

  myuuid:=uuid.New()
  fmt.Println(myuuid)
}

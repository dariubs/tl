package main

import (
        "fmt"
        "os"
        )

var (
  file  *os.File
  err   error
  filename string = "todo"
)


func main(){

  args := os.Args[1:]

  if args[0] == "create" {
    file, err = os.Create(filename)
    return
  } else {
    file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
  }


  if err != nil {
    panic(err)
  }

  defer func(){
    if err := file.Close(); err != nil {
      panic(err)
    }
  }()


  text := args[0] + "\n"
  _, err = file.WriteString(text)
  err = file.Sync()

  fmt.Printf("added")
}

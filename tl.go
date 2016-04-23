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

  if len(os.Args) < 2 {
    fmt.Printf("Wrong number of arguments")
    os.Exit(1)
  }

  args := os.Args[1:]


  if args[0] == "create" {
    file, err = os.Create(filename)
    return
  } else {
    file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
  }


  if err != nil {
    fmt.Printf("an error occured")
  }

  defer func(){
    if err := file.Close(); err != nil {
      fmt.Printf("an error occured to close file")
    }
  }()


  text := args[0] + "\n"
  _, err = file.WriteString(text)
  err = file.Sync()

  fmt.Printf("added")
}

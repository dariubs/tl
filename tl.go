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
    _, err := os.Stat(filename)

    if os.IsNotExist(err){
        file, err = os.Create(filename)
        fmt.Println("todo file created.")
    } else {
      fmt.Println("todo file already exists.")
    }

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

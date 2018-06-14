package main

import (
  "fmt"
  "log"
  "os"
  "strconv"

  "github.com/xanderflood/go-tropical/pkg/bisplit"
)

func main() {
  if len(os.Args) != 3 {
    log.Fatal("enum-bisplits require exactly two integer command-line arguments")
  }

  w64, err := strconv.ParseUint(os.Args[1], 10, 64)
  if err != nil {
    log.Fatal(err)
  }
  b64, err := strconv.ParseUint(os.Args[2], 10, 64)
  if err != nil {
    log.Fatal(err)
  }

  w := uint(w64)
  b := uint(b64)

  enum := bisplit.Enumerate(w, b)
  sp := bisplit.ModuliSpace{W: w, B: b}

  for bs := range enum {
    fmt.Println(sp.String(bs))
  }
}

package main

import "fmt"

import "github.com/ianohf/go-internal-pkg/foo"

func main() {
  fmt.Println("in main")

  foo.PrFoo("calling from main")
}


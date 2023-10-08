package main

import "fmt"

import "github.com/ianohf/go-internal-pkg/pkg"
import "github.com/ianohf/go-internal-pkg/pkg/foo"
import "github.com/ianohf/go-internal-pkg/pkg/bar"
//import "github.com/ianohf/go-internal-pkg/pkg/foo/internal"
import "github.com/ianohf/go-internal-pkg/pkg/foo/sibling"

func main() {
  fmt.Println("in main")

  pkg.PrExample("calling from main")
  foo.PrFoo("calling from main")
  bar.PrBar("calling from main")

  sibling.PrSibling("calling from main")
  // internal.PrInternal("calling from main")

  foo.PrintInternal()
  sibling.PrintInternal()
}


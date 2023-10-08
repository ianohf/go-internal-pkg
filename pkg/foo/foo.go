package foo

import "fmt"
import "github.com/ianohf/go-internal-pkg/pkg/foo/internal"

func PrFoo(s string) {
  fmt.Printf("print from foo: %s\n", s)
}

func PrintInternal() {
  internal.PrInternal("calling from foo")
}


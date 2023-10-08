package sibling

import "fmt"
import "github.com/ianohf/go-internal-pkg/pkg/foo/internal"

func PrSibling(s string) {
  fmt.Printf("print from sibling: %s\n", s)
}

func PrintInternal() {
  internal.PrInternal("calling from sibling")
}


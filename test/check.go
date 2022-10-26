package test

import (
	"encoding/json"
	"strings"

	g "github.com/onsi/gomega"
)

func Check(i int, expect, actual, input any) {
	sb := &strings.Builder{}
	enc := json.NewEncoder(sb)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(input)
	g.Expect(actual).WithOffset(1).To(g.Equal(expect), "test case %d with input %s", i, sb.String())
}

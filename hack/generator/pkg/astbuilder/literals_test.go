/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_LiteralString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		original string
		expected string
	}{
		{"a word", "foo", "\"foo\""},
		{"multiple words", "foo bar baz", "\"foo bar baz\""},
		{"existing quotes escaped", "foo \"bar\" baz", "\"foo \\\"bar\\\" baz\""},
		{"existing slashes escaped", "this/or/that this\\or\\that", "\"this/or/that this\\\\or\\\\that\""},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			literal := StringLiteral(c.original)

			g.Expect(literal.Value).To(Equal(c.expected))
		})
	}
}

package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6h2v4.079l3.341-2.34l1.147 1.639L13.743 12l3.745 2.622l-1.147 1.639L13 13.92V18h-2v-4.079l-3.341 2.34l-1.148-1.639L10.257 12L6.51 9.378l1.15-1.639L11 10.08V6Z"/>`),
		g.Group(children),
	)
}
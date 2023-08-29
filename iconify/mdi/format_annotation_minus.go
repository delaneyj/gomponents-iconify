package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAnnotationMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 7h-2L3 21h2.2l1.1-3h6.2l1.1 3H16L10.5 7m-3.4 9l2.4-6.3l2.4 6.3H7.1M22 7h-8V5h8v2Z"/>`),
		g.Group(children),
	)
}
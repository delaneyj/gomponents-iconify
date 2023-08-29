package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.41 7.41L10 12l-4.59 4.59L4 15.17L7.17 12L4 8.83l1.41-1.42m13.18 9.18L14 12l4.59-4.58L20 8.83L16.83 12L20 15.17l-1.41 1.42Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.76 16.83L14.59 14l-2.83-2.83l1.41-1.41L16 12.59l2.83-2.83l1.41 1.41L17.41 14l2.83 2.83l-1.41 1.41L16 15.41l-2.83 2.83l-1.41-1.41M2 11h3l2.29 5.4L10 6h12v2H11.55L8.68 19H6.22l-2.54-6H2v-2Z"/>`),
		g.Group(children),
	)
}
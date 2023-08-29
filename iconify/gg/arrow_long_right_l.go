package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongRightL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.916 7.757l4.25 4.236l-4.236 4.25l-1.416-1.412l1.819-1.825l-16.5.022v2.002h-2v-6h2v1.998l16.51-.022l-1.838-1.832l1.411-1.417Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.243 7.172l-1.415-1.415l-9.07 9.071v-4.585h-2v8h8v-2H9.17l9.072-9.071Z"/>`),
		g.Group(children),
	)
}
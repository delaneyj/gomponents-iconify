package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.757 7.172l1.415-1.415l9.07 9.071v-4.585h2v8h-8v-2h4.586l-9.07-9.071Z"/>`),
		g.Group(children),
	)
}
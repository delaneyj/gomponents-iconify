package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h3v-3H3m5 3h3v-7H8m5 7h3V9h-3m5 12h3V3h-3v18Z"/>`),
		g.Group(children),
	)
}
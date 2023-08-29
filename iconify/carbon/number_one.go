package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 10v12v-12m1-1h-5v2h3v10h-3v2h8v-2h-3V9Z"/>`),
		g.Group(children),
	)
}
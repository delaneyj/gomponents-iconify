package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.346 5.353c.21-.129.428-.246.654-.353a2 2 0 1 1 4 0a7 7 0 0 1 4 6v3m-1 3H4a4 4 0 0 0 2-3v-3a6.996 6.996 0 0 1 1.273-3.707M9 17v1a3 3 0 0 0 6 0v-1M3 3l18 18"/>`),
		g.Group(children),
	)
}
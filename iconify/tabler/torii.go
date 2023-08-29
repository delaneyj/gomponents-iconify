package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torii(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4c5.333 1.333 10.667 1.333 16 0M4 8h16m-8-3v3m6-3.5V20M6 4.5V20"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 18v-5l4-2l4 2v5l-4 2zm4-7V6l4-2l4 2v5m-4 2l4-2l4 2v5l-4 2l-4-2"/>`),
		g.Group(children),
	)
}
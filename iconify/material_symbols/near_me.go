package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.9 21l-2.85-7.05L3 11.1V9.7L21 3l-6.7 18h-1.4Z"/>`),
		g.Group(children),
	)
}
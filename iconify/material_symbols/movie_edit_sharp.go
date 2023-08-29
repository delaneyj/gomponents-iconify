package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieEditSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V3h2l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h5v6H4v8h8v2H2Zm17.15-6.3l2.15 2.1l-5.175 5.2H14v-2.125l5.15-5.175Zm2.875 1.4L19.9 11.975l1.425-1.425l2.1 2.15l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}
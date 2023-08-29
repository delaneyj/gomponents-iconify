package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallLampSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v6H3Zm3-2v-2h6v-4H5.65l3-10h8.7l3 10H14v6H6Z"/>`),
		g.Group(children),
	)
}
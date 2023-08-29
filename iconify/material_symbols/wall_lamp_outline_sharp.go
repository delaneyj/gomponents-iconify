package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallLampOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v6H3Zm5.35-10h9.3l-1.8-6h-5.7l-1.8 6Zm0 0h9.3h-9.3ZM6 19v-2h6v-4H5.65l3-10h8.7l3 10H14v6H6Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesearchRollerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23v-8h2v-3H2V4h4V2h14v6H6V6H4v4h10v5h2v8h-6ZM8 4v2v-2Zm4 17h2h-2Zm0 0h2v-4h-2v4ZM8 6h10V4H8v2Z"/>`),
		g.Group(children),
	)
}
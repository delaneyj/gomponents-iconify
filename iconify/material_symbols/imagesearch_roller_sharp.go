package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesearchRollerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23v-8h2v-3H2V4h4V2h14v6H6V6H4v4h10v5h2v8h-6Z"/>`),
		g.Group(children),
	)
}
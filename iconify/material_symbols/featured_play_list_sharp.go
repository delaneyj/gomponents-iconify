package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeaturedPlayListSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 13h9v-2H6v2Zm0-3h9V8H6v2ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}
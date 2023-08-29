package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAspectRatioSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 12h2v-2H6v2Zm4 0h2v-2h-2v2Zm4 4h2v-2h-2v2Zm0-4h2v-2h-2v2ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}
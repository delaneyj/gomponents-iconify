package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiveTvSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 15.5l7-4.5l-7-4.5v9ZM8 21v-2H2V3h20v16h-6v2H8Z"/>`),
		g.Group(children),
	)
}
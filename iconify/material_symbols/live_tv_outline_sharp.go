package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiveTvOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 15.5l7-4.5l-7-4.5v9ZM8 21v-2H2V3h20v16h-6v2H8Zm-4-4h16V5H4v12Zm0 0V5v12Z"/>`),
		g.Group(children),
	)
}
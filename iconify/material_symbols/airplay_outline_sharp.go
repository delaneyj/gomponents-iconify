package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 21l6-6l6 6H6Zm-4-2V3h20v16h-5v-2h3V5H4v12h3v2H2Zm10-7Z"/>`),
		g.Group(children),
	)
}
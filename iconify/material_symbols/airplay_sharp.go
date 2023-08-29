package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 21l6-6l6 6H6Zm-4-2V3h20v16h-3l-7-7l-7 7H2Z"/>`),
		g.Group(children),
	)
}
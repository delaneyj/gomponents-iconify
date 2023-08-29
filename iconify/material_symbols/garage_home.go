package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V9l8-6l8 6v12h-3V11H7v10H4Zm5-2h6v-2H9v2Zm0-4h6v-2H9v2Z"/>`),
		g.Group(children),
	)
}
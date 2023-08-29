package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lists(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-4h4v4H2Zm6 0v-4h14v4H8Zm-6-6v-4h4v4H2Zm6 0v-4h14v4H8ZM2 8V4h4v4H2Zm6 0V4h14v4H8Z"/>`),
		g.Group(children),
	)
}
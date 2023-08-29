package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSharpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-6q0-.825.588-1.413T8 13h8V6.8l-1.6 1.6L13 7l4-4l4 4l-1.4 1.4L18 6.8V13q0 .825-.588 1.413T16 15H8v6H6Z"/>`),
		g.Group(children),
	)
}
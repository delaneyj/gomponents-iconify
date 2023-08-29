package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTurnSharpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6.83L4.41 8.41L3 7l4-4l4 4l-1.41 1.41L8 6.83V13h10v8h-2v-6H6z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.8 21l-5.2-9l5.2-9h10.4l5.2 9l-5.2 9H6.8Z"/>`),
		g.Group(children),
	)
}
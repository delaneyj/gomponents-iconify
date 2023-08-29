package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func North(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22V5.825L6.4 10.4L5 9l7-7l7 7l-1.4 1.425l-4.6-4.6V22h-2Z"/>`),
		g.Group(children),
	)
}
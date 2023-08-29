package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-6-6l1.425-1.425L12 19.15l4.575-4.575L18 16l-6 6ZM7.45 9.4L6 8l6-6l6 6l-1.45 1.4L12 4.85L7.45 9.4Z"/>`),
		g.Group(children),
	)
}
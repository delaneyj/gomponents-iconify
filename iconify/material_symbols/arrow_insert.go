package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.6 18L8 8.4V17H6V5h12v2H9.4l9.6 9.6l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}
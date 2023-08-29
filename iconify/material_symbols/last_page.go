package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 18l-1.4-1.4l4.6-4.6l-4.6-4.6L7 6l6 6l-6 6Zm9 0V6h2v12h-2Z"/>`),
		g.Group(children),
	)
}
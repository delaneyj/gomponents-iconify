package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Details(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21L12 3l10 18H2Zm3.4-2H11V8.925L5.4 19Zm7.6 0h5.6L13 8.925V19Z"/>`),
		g.Group(children),
	)
}
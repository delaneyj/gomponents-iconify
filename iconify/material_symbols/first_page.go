package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V6h2v12H6Zm11 0l-6-6l6-6l1.4 1.4l-4.6 4.6l4.6 4.6L17 18Z"/>`),
		g.Group(children),
	)
}
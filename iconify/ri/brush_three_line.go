package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19.997v-5h2v5h9v-7H5v7h3Zm-4-9h16v-3h-6v-4h-4v4H4v3Zm-1 10v-8H2v-6a1 1 0 0 1 1-1h5v-3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3h5a1 1 0 0 1 1 1v6h-1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}
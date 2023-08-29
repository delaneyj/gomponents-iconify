package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm14 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm0 14a3 3 0 1 1 0-6a3 3 0 0 1 0 6ZM5 22a3 3 0 1 1 0-6a3 3 0 0 1 0 6ZM9 4h6v2H9V4Zm0 14h6v2H9v-2ZM4 9h2v6H4V9Zm14 0h2v6h-2V9Z"/>`),
		g.Group(children),
	)
}
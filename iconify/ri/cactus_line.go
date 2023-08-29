package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.998 2a4 4 0 0 1 4 4v9h1c.55 0 1-.45 1-1V8a1 1 0 0 1 2 0v6c0 1.66-1.34 3-3 3h-1v3h2v2h-12v-2h2v-6h-1a3 3 0 0 1-3-3V9a1 1 0 1 1 2 0v2c0 .55.45 1 1 1h1V6a4 4 0 0 1 4-4Zm0 2a2 2 0 0 0-2 2v14h4V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
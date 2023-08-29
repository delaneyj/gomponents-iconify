package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6h-1a3 3 0 0 0-3 3v8a1 1 0 0 0 2 0v-4h3v4a1 1 0 0 0 2 0V9a3 3 0 0 0-3-3Zm1 5h-3V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1ZM8 6a1 1 0 0 0-1 1v5.76L3.89 6.55A1 1 0 0 0 2 7v10a1 1 0 0 0 2 0v-5.76l3.11 6.21A1 1 0 0 0 8 18a.91.91 0 0 0 .23 0A1 1 0 0 0 9 17V7a1 1 0 0 0-1-1Zm4-2a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
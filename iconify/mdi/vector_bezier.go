package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorBezier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 4A1.5 1.5 0 0 0 6 5.5A1.5 1.5 0 0 0 7.5 7c.63 0 1.2-.4 1.41-1H13c.67-.67 1.33-1 2-1H8.91c-.21-.6-.78-1-1.41-1M19 5C8 5 14 17 5 17v2c11 0 5-12 14-12V5m-2.5 12c-.63 0-1.2.4-1.41 1H11c-.67.67-1.33 1-2 1h6.09c.21.6.78 1 1.41 1a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}
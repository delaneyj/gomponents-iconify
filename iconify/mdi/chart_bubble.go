package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.2 11.2c1.77 0 3.2 1.43 3.2 3.2c0 1.77-1.43 3.2-3.2 3.2c-1.77 0-3.2-1.43-3.2-3.2c0-1.77 1.43-3.2 3.2-3.2m7.6 4.8a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2m.4-12A4.8 4.8 0 0 1 20 8.8c0 2.65-2.15 4.8-4.8 4.8a4.8 4.8 0 0 1-4.8-4.8c0-2.65 2.15-4.8 4.8-4.8Z"/>`),
		g.Group(children),
	)
}
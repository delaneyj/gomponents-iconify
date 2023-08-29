package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17v2h-8v-2h8M8 14c0-1.77 2-5.04 4-7.61c1.5 1.94 3 4.28 3.67 6.08c.6-.26 1.26-.41 1.95-.47C16.31 8.1 12 3.25 12 3.25S6 10 6 14c0 3.31 2.69 6 6 6h.34c-.22-.64-.34-1.3-.34-2c-2.21 0-4-1.79-4-4Z"/>`),
		g.Group(children),
	)
}
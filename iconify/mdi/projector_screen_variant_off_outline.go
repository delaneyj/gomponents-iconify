package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenVariantOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l3 3H4c-.55 0-1 .45-1 1v1c0 .55.45 1 1 1h1v9h11.11l4.73 4.73l1.27-1.27M7 16V9h.11l7 7H7m5.2-7l-3-3H20c.55 0 1 .45 1 1v1c0 .55-.45 1-1 1h-1v6.8l-2-2V9h-4.8Z"/>`),
		g.Group(children),
	)
}
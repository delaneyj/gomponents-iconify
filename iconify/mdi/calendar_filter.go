package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 1h2v2h8V1h2v2h1a2 2 0 0 1 2 2v6h-2V8H5v11h10v2H5a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h1V1m11 20l1.8 1.77c.5.5 1.2.1 1.2-.49V18l2.8-3.4A1 1 0 0 0 22 13h-7c-.8 0-1.3 1-.8 1.6L17 18v3"/>`),
		g.Group(children),
	)
}
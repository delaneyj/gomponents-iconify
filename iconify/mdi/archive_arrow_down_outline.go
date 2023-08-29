package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveArrowDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H4V10h2v9h12v-9h2v11M3 3h18v6H3V3m2 2v2h14V5m-8.5 6v3H8l4 4l4-4h-2.5v-3"/>`),
		g.Group(children),
	)
}
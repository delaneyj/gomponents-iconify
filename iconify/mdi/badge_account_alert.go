package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeAccountAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3h-3v3H8V3H5a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m-5 5a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2m4 8H6v-1c0-1.33 2.67-2 4-2s4 .67 4 2v1M11 5H9V1h2v4m3 14H6v-1h8v1m-4 2H6v-1h4v1m9-8V7h2v6h-2m0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}
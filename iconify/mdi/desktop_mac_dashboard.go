package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMacDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14V4H3v10h18m0-12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-7l2 3v1H8v-1l2-3H3a2 2 0 0 1-2-2V4c0-1.11.89-2 2-2h18M4 5h11v5H4V5m12 0h4v2h-4V5m4 3v5h-4V8h4M4 11h5v2H4v-2m6 0h5v2h-5v-2Z"/>`),
		g.Group(children),
	)
}
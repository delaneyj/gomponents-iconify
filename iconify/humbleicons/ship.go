package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l.756.378a3 3 0 0 0 2.523.074l1.04-.446a3 3 0 0 1 2.363 0l1.04.446a3 3 0 0 0 2.523-.074l.413-.207a3 3 0 0 1 2.684 0l.547.273a3 3 0 0 0 2.29.163L21 17M5 14.5L4 10h4m10 4.5l2.5-4.5h-8.245H13.5m0 0l-.721-3H8v3m5.5 0H8m3 3h.1M10 4.5l-.2-.2a2 2 0 0 0-1.899-.525l-.336.084a2 2 0 0 1-1.118-.043L5.5 3.5"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11h14V8H5zm12.5 0L16 21H8L6.5 11M6 8V7a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v1m-3-3V3"/>`),
		g.Group(children),
	)
}
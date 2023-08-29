package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsTransferDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3v6m-7 9l-3 3l-3-3m3 3V3m13 3l-3-3l-3 3m3 15v-2m0-4v-2"/>`),
		g.Group(children),
	)
}
package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2M7 17v2a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-2M7 17H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1m0 0V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2M6 7h12m0 0h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2M7 10h1"/>`),
		g.Group(children),
	)
}
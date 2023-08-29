package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataRefineryReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm24-8h-6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zm-6-8v6h6V4z"/><path fill="currentColor" d="M24 26h-8v-2h8v-7H8a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h10v2H8v7h16a2.002 2.002 0 0 1 2 2v7a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}
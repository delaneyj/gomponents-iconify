package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 24H10a2 2 0 0 1-2-2v-3h2v3h12v-3h2v3a2 2 0 0 1-2 2zM2 15h28v2H2zm22-2h-2v-3H10v3H8v-3a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm6-3h-2V4h-6V2h8v8zM4 10H2V2h8v2H4v6zm6 20H2v-8h2v6h6v2zm20 0h-8v-2h6v-6h2v8z"/>`),
		g.Group(children),
	)
}
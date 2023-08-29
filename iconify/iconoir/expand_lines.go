package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2H6m12 20H6m6-8v5m0 0l3-3m-3 3l-3-3m3-6V5m0 0l3 3m-3-3L9 8"/>`),
		g.Group(children),
	)
}
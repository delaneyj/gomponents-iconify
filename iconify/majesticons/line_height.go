package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10V5m0 0L4 7m2-2l2 2m-2 7v5m0 0l2-2m-2 2l-2-2m8-10h8m0 5h-8m0 5h8"/>`),
		g.Group(children),
	)
}
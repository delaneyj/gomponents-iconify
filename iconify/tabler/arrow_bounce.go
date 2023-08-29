package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBounce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 18h4M3 8a9 9 0 0 1 9 9v1l1.428-4.285a12 12 0 0 1 6.018-6.938L20 6.5"/><path d="M15 6h5v5"/></g>`),
		g.Group(children),
	)
}
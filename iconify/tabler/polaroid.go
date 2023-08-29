package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polaroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 10h16"/><path d="m4 12l3-3c.928-.893 2.072-.893 3 0l4 4"/><path d="m13 12l2-2c.928-.893 2.072-.893 3 0l2 2m-6-5h.01"/></g>`),
		g.Group(children),
	)
}
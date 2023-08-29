package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3.707 6.293l2.586-2.586a1 1 0 0 1 1.414 0l5.586 5.586a1 1 0 0 1 0 1.414l-2.586 2.586a1 1 0 0 1-1.414 0L3.707 7.707a1 1 0 0 1 0-1.414z"/><path d="m6 10l-3 3l3 3l3-3m1-7l3-3l3 3l-3 3m-1 3l1.5 1.5m1 3.5a2.5 2.5 0 0 0 2.5-2.5M15 21a6 6 0 0 0 6-6"/></g>`),
		g.Group(children),
	)
}
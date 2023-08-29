package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10V7.5L14.5 2H6a2 2 0 0 0-2 2v16c0 1.1.9 2 2 2h4.5"/><path d="M14 2v6h6m-4 14a2 2 0 0 1-2-2m6 2a2 2 0 0 0 2-2m-2-6a2 2 0 0 1 2 2m-6-2a2 2 0 0 0-2 2"/></g>`),
		g.Group(children),
	)
}
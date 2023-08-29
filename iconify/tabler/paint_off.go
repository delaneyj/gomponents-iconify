package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3h10a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-4M9 9H7a2 2 0 0 1-2-2V5"/><path d="M19 6h1a2 2 0 0 1 2 2a5 5 0 0 1-5 5m-4 0h-1v2m-2 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1zM3 3l18 18"/></g>`),
		g.Group(children),
	)
}
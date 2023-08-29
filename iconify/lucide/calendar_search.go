package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7.5M16 2v4M8 2v4m-5 4h18"/><path d="M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6v0Zm4 1l-1.5-1.5"/></g>`),
		g.Group(children),
	)
}
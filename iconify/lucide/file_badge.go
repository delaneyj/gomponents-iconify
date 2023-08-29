package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 7V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-6"/><path d="M14 2v6h6M5 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M7 16.5L8 22l-3-1l-3 1l1-5.5"/></g>`),
		g.Group(children),
	)
}
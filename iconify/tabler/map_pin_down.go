package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M12.736 21.345a2 2 0 0 1-2.149-.445l-4.244-4.243a8 8 0 1 1 13.59-4.624M19 16v6m3-3l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}
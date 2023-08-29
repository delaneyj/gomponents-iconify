package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 12l-8.5 8.5a2.12 2.12 0 1 1-3-3L11 9"/><path d="M15 13L9 7l4-4l6 6h3a8 8 0 0 1-7 7z"/></g>`),
		g.Group(children),
	)
}
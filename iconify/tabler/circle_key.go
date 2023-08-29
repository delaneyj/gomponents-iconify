package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 10a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="m12.5 11.5l-4 4L10 17m2-2l-1.5-1.5"/></g>`),
		g.Group(children),
	)
}
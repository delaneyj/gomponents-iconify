package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h18L14.079 5.388a2.3 2.3 0 0 0-4.158 0L3 20z"/><path d="m7.5 11l2 2.5L12 11l2 3l2.5-2"/></g>`),
		g.Group(children),
	)
}
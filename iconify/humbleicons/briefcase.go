package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 8V6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M4 14l3.15.787a20 20 0 0 0 9.7 0L20 14"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8z"/><circle cx="12" cy="12" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
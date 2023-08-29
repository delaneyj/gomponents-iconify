package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 5h12l3 5l-8.5 9.5a.7.7 0 0 1-1 0L3 10l3-5"/><path d="M10 12L8 9.8l.6-1"/></g>`),
		g.Group(children),
	)
}
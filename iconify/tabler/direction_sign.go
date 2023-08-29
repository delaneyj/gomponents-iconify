package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3.32 12.774l7.906 7.905c.427.428 1.12.428 1.548 0l7.905-7.905a1.095 1.095 0 0 0 0-1.548l-7.905-7.905a1.095 1.095 0 0 0-1.548 0l-7.905 7.905a1.095 1.095 0 0 0 0 1.548zM8 12h7.5"/><path d="m12 8.5l3.5 3.5l-3.5 3.5"/></g>`),
		g.Group(children),
	)
}
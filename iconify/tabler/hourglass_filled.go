package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 2a2 2 0 0 1 1.995 1.85L19 4v2a6.996 6.996 0 0 1-3.393 6a6.994 6.994 0 0 1 3.388 5.728L19 18v2a2 2 0 0 1-1.85 1.995L17 22H7a2 2 0 0 1-1.995-1.85L5 20v-2a6.996 6.996 0 0 1 3.393-6a6.994 6.994 0 0 1-3.388-5.728L5 6V4a2 2 0 0 1 1.85-1.995L7 2h10z"/></g>`),
		g.Group(children),
	)
}
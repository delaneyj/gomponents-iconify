package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDivisionSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4Z"/><path stroke-width="3.75" d="M12 16h.01v.01H12zm0-8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M15 12H9"/></g>`),
		g.Group(children),
	)
}
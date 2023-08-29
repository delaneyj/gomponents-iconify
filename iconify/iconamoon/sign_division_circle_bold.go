package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDivisionCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M15 12H9"/></g>`),
		g.Group(children),
	)
}
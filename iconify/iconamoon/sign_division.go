package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDivision(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M18 12H6"/></g>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignPercentBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="16" cy="18" r="2"/><circle cx="8" cy="6" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/></g>`),
		g.Group(children),
	)
}
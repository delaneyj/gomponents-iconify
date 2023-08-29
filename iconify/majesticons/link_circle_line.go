package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="9"/><path d="M10 9c-1 0-3 .6-3 3s2 3 3 3m4-6c1 0 3 .6 3 3s-2 3-3 3m-4-3h4"/></g>`),
		g.Group(children),
	)
}
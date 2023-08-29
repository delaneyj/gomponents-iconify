package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M11.87 21.48a1.992 1.992 0 0 1-1.283-.58l-4.244-4.243a8 8 0 1 1 13.355-3.474M15 19l2 2l4-4"/></g>`),
		g.Group(children),
	)
}
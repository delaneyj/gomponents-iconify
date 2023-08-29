package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReloadCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m13 8l2 2l-2 2m0-2h-1c-1 0-3 .6-3 3s2 3 3 3c.534 0 1.353-.171 2-.695"/></g>`),
		g.Group(children),
	)
}
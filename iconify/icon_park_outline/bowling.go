package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M24.777 44c11.05 0 20-8.95 20-20s-8.95-20-20-20s-20 8.95-20 20s8.95 20 20 20Z"/><path fill="currentColor" d="M24.777 26c1.66 0 3-1.34 3-3s-1.34-3-3-3s-3 1.34-3 3s1.34 3 3 3Zm7-8c1.66 0 3-1.34 3-3s-1.34-3-3-3s-3 1.34-3 3s1.34 3 3 3Zm-14 0c1.66 0 3-1.34 3-3s-1.34-3-3-3s-3 1.34-3 3s1.34 3 3 3Z"/></g>`),
		g.Group(children),
	)
}
package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnionAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 22A7 7 0 1 0 9 8a7 7 0 0 0 0 14Z"/><path d="M15 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/></g>`),
		g.Group(children),
	)
}
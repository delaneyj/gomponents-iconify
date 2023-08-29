package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 7l2 13h14l2-13l-5 3l-4-6l-4 6l-5-3z"/><circle cx="12" cy="14" r="2"/></g>`),
		g.Group(children),
	)
}
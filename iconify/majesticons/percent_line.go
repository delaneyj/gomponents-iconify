package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18L18 6"/><circle cx="16.5" cy="16.5" r="2.5"/><circle cx="7.5" cy="7.5" r="2.5"/></g>`),
		g.Group(children),
	)
}
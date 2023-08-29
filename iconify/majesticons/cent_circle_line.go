package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M12 16a4 4 0 0 1 0-8m0 8v2m0-2a3.987 3.987 0 0 0 2.828-1.172M12 8V6m0 2c1.105 0 2.105.448 2.828 1.172"/></g>`),
		g.Group(children),
	)
}
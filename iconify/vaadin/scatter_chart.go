package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M5 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm6-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-3 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}
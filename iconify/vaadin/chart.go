package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 15h16v1H0v-1z"/><path fill="currentColor" d="M0 0h1v16H0V0zm9 8L6.1 5L2 9v5h14V.9z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 0v7c0 5.6 7 9 7 9s7-3.4 7-9V0H1zm13 7c0 4.2-4.6 7.1-6 7.9V1h6v6z"/>`),
		g.Group(children),
	)
}
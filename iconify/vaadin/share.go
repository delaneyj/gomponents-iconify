package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 3H4.9S0 3 0 8c0 3.9 3 8 3 8S1.3 9 4.8 9H10v3l6-6l-6-6v3z"/>`),
		g.Group(children),
	)
}
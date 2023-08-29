package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 7.9L10 3v3H8c-8 0-8 8-8 8s1-4 7.8-4H10v2.9l6-5z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2v12h16V2H0zm13 7H6v2L3 8l3-3v2h7v2z"/>`),
		g.Group(children),
	)
}
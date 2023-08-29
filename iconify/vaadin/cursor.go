package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 0v13l3.31-3.47L10 16l1.37-.63L8.65 9H13L4 0z"/>`),
		g.Group(children),
	)
}
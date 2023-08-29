package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccordionMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 4v8h16V4H0zm15 7H1V7h14v4zM0 0h16v3H0V0zm0 13h16v3H0v-3z"/>`),
		g.Group(children),
	)
}
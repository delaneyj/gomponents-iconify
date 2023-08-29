package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Institution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0L0 3v2h16V3zM0 14h16v2H0v-2zm16-7V6H0v1h1v5H0v1h16v-1h-1V7h1zM4 12H3V7h1v5zm3 0H6V7h1v5zm3 0H9V7h1v5zm3 0h-1V7h1v5z"/>`),
		g.Group(children),
	)
}
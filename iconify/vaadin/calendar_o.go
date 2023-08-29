package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1h-2zm1 14H1V6h14v9z"/><path fill="currentColor" d="M3 0h1v3H3V0zm9 0h1v3h-1V0z"/>`),
		g.Group(children),
	)
}
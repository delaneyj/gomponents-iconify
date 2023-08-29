package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restaurant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.5 24V.5h.5A4.5 4.5 0 0 1 20.5 5v9.5H20a2.5 2.5 0 0 0-2.5 2.5M7 24V9.5m0 0A3.5 3.5 0 0 0 10.5 6V0M7 9.5A3.5 3.5 0 0 1 3.5 6V0M7 0v7.5"/>`),
		g.Group(children),
	)
}
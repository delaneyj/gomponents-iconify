package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotVacuumVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3a2 2 0 0 0-2 2v2h2V5h14v2h2V5a2 2 0 0 0-2-2H5m3 4v2h8V7H8M3 9v3a9 9 0 0 0 9 9a9 9 0 0 0 9-9V9h-2v3a7 7 0 0 1-7 7a7 7 0 0 1-7-7V9H3m9 3a2.5 2.5 0 0 0-2.5 2.5A2.5 2.5 0 0 0 12 17a2.5 2.5 0 0 0 2.5-2.5A2.5 2.5 0 0 0 12 12Z"/>`),
		g.Group(children),
	)
}
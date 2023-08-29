package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScheduleNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M14 0v3M1 7h18M6 0v3m-2 8h2m2 0h8M4 15h2m2 0h6m-2 4H1V3h18v10m0 2v9m-4-7l8 5m0-5l-8 5"/>`),
		g.Group(children),
	)
}
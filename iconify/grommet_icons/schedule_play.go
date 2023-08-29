package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchedulePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M14 0v3M1 7h18M6 0v3m-2 8h2m2 0h8M4 15h2m2 0h6m-1 4H1V3h18v10m-1 10a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm-.5-6l1.5 1l-1.5 1v-2Z"/>`),
		g.Group(children),
	)
}
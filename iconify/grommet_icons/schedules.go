package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17 7h6v16H7v-4m16-8h-6M13 0v3M1 7h16M1 3h16v16H1V3Zm4-3v3m-1 8h2m2 0h6M4 15h2m2 0h6"/>`),
		g.Group(children),
	)
}
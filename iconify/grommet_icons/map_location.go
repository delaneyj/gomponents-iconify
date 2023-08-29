package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17.5 6.5L23 9v13l-7-3l-8 3l-7-3V6l5 2m10 11v-7M8 22V12m4 4.273S6 11.5 6 7c0-3.75 3-6 6-6s6 2.25 6 6c0 4.5-6 9.273-6 9.273ZM13 7a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighSpeedRail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 20v10a8 8 0 0 0 8 8h14a8 8 0 0 0 8-8V20m-5 18v4m-20-4v4m6 2h8"/><path stroke-linejoin="round" stroke-width="4" d="M9 16.36C9 13 15 4 24 4s15 9 15 12.36V20H9v-3.64Z"/><path fill="currentColor" d="M20.5 32a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 11h4"/></g>`),
		g.Group(children),
	)
}
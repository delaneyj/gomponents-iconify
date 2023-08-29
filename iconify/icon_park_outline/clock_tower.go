package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 44h40M27 32h12v12H27zm11-22v6m-10-6v6m0-6l5-6l5 6H28Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M25 20H11a2 2 0 0 0-2 2v22"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M15 29h4m-4 7h4"/><rect width="16" height="16" x="25" y="16" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="33" cy="24" r="3" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M33 32v10"/></g>`),
		g.Group(children),
	)
}
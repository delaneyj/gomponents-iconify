package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClockTower0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 44h40M27 32h12v12H27zm11-22v6m-10-6v6m0-6l5-6l5 6H28Z"/><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M25 20H11a2 2 0 0 0-2 2v22"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M15 29h4m-4 7h4"/><rect width="16" height="16" x="25" y="16" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="33" cy="24" r="3" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M33 32v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClockTower0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RockingHorse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRockingHorse0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M14 29s-3-5-3-11h16s0-4.5 3-8s6-5 6-5l8 8.5v5L36 16c-4 5-2 13-2 13H14Z"/><path d="m30 29l5 12M18 29l-5 12m-9-4s6 6 20 6s20-6 20-6M11 18c0-3-2-6-7-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRockingHorse0)"/>`),
		g.Group(children),
	)
}
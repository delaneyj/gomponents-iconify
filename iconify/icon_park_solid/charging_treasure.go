package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingTreasure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChargingTreasure0"><g fill="none"><path stroke="#fff" stroke-width="4" d="M9.975 8.557A3 3 0 0 1 12.942 6h23.036a3 3 0 0 1 2.979 2.646l3.145 26.5a3 3 0 0 1-2.98 3.354H8.983a3 3 0 0 1-2.967-3.443l3.96-26.5Z"/><rect width="36" height="12" x="6" y="30" fill="#fff" stroke="#fff" stroke-width="4" rx="6"/><rect width="4" height="4" x="19" y="34" fill="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 36h6"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22 12h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChargingTreasure0)"/>`),
		g.Group(children),
	)
}
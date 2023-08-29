package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicLocksClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectronicLocksClose0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="24" height="18" x="12" y="20" fill="#555" rx="2"/><path stroke-linecap="round" d="M18 20v-6c0-3.682 2.686-6 6-6s6 2.318 6 6v6m-6 8v2M6 18v12m36-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectronicLocksClose0)"/>`),
		g.Group(children),
	)
}
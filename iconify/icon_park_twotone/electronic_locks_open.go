package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicLocksOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectronicLocksOpen0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="24" height="18" x="12" y="20" fill="#555" rx="2"/><path stroke-linecap="round" d="M18 20v-6c0-3.682 2.686-6 6-6c1.85 0 3.503.722 4.604 2a5.62 5.62 0 0 1 1.102 2M24 28v2M6 18v12m36-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectronicLocksOpen0)"/>`),
		g.Group(children),
	)
}
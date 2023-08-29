package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectricWave0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path d="M12 25h3l4-11l3 22l5-13l4 6l3-4h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectricWave0)"/>`),
		g.Group(children),
	)
}
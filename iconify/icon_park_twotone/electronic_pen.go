package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectronicPen0"><g fill="none"><rect width="12" height="38" x="35.193" y="5.322" fill="#555" stroke="#fff" stroke-width="4" rx="6" transform="rotate(45 35.193 5.322)"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m18 23l8 8M6 43l6-6"/><rect width="4" height="4" x="33.268" y="12.34" fill="#fff" rx="2" transform="rotate(30 33.268 12.34)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectronicPen0)"/>`),
		g.Group(children),
	)
}
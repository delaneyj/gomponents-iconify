package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLaptopComputer0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="38" height="24" x="5" y="8" fill="#555" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 40h40M22 14h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLaptopComputer0)"/>`),
		g.Group(children),
	)
}
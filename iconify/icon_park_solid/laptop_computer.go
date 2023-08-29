package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLaptopComputer0"><g fill="none" stroke-width="4"><rect width="38" height="24" x="5" y="8" fill="#fff" stroke="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M4 40h40"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M22 14h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLaptopComputer0)"/>`),
		g.Group(children),
	)
}
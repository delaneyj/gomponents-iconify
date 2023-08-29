package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCylinder0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 10c0 3.314-8.954 6-20 6S4 13.314 4 10s8.954-6 20-6s20 2.686 20 6Zm0 28c0 3.314-8.954 6-20 6S4 41.314 4 38s8.954-6 20-6s20 2.686 20 6Z"/><path d="M44 10v28M4 10v28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCylinder0)"/>`),
		g.Group(children),
	)
}
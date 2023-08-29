package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTaurus0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="31" r="9" fill="#fff"/><path stroke-linecap="round" d="M44 8c0 7.732-8.954 14-20 14S4 15.732 4 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTaurus0)"/>`),
		g.Group(children),
	)
}
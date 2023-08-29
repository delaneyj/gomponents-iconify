package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Family(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFamily0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M10 19s-5.143 2-6 9m34-9s5.143 2 6 9m-26-9s4.8 1.167 6 7m6-7s-4.8 1.167-6 7m-4 8s-4.2.75-6 6m14-6s4.2.75 6 6"/><circle cx="24" cy="31" r="5" fill="#555" stroke-linejoin="round"/><circle cx="34" cy="14" r="6" fill="#555" stroke-linejoin="round"/><circle cx="14" cy="14" r="6" fill="#555" stroke-linejoin="round"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFamily0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Division(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDivision0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="5" fill="#555"/><circle cx="24" cy="37" r="5" fill="#555"/><path d="M44 24H4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDivision0)"/>`),
		g.Group(children),
	)
}
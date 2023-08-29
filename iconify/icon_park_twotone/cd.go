package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCd0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="18"/><path stroke-linecap="round" d="M13 24c0-6.075 4.925-11 11-11"/><circle cx="24" cy="24" r="5" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCd0)"/>`),
		g.Group(children),
	)
}
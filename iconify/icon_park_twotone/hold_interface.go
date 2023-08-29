package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoldInterface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHoldInterface0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m20 33l6 2s15-3 17-3s2 2 0 4s-9 8-15 8s-10-3-14-3H4"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 29c2-2 6-5 10-5s13.5 4 15 6s-3 5-3 5"/><rect width="16" height="6" x="26" y="15" fill="#555" rx="3"/><path stroke-linecap="round" d="M26 9h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHoldInterface0)"/>`),
		g.Group(children),
	)
}
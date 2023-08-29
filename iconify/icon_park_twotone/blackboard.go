package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlackboard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 7h32v24H8z"/><path stroke-linecap="round" d="M4 7h40M15 41l9-10l9 10M16 13h16m-16 6h12m-12 6h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlackboard0)"/>`),
		g.Group(children),
	)
}
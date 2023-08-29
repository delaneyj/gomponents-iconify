package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlackboard0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 7h32v24H8z"/><path stroke="#fff" stroke-linecap="round" d="M4 7h40M15 41l9-10l9 10"/><path stroke="#000" stroke-linecap="round" d="M16 13h16m-16 6h12m-12 6h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlackboard0)"/>`),
		g.Group(children),
	)
}
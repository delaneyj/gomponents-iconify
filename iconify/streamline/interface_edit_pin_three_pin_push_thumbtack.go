package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPinThreePinPushThumbtack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="5" r="4.5"/><path d="M10.25 3.67a1.5 1.5 0 0 0-2.31-.23M.5 13.5l5.58-5.07"/></g>`),
		g.Group(children),
	)
}
package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditPinOnePinPushThumbtack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="5" r="4.5"/><path d="m.5 13.5l5.58-5.07"/><circle cx="8.5" cy="3.5" r=".5"/></g>`),
		g.Group(children),
	)
}
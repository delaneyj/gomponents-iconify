package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerLoadingThreeProgressLoadingDotLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="1.5"/><circle cx="12.25" cy="7" r="1.25"/><circle cx="1.75" cy="7" r="1.25"/></g>`),
		g.Group(children),
	)
}
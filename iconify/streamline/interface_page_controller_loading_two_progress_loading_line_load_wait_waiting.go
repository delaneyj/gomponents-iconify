package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerLoadingTwoProgressLoadingLineLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5" height="4" x=".5" y="5" rx=".5"/><rect width="3" height="4" x="8" y="5" rx=".5"/><path d="M13.5 5v4"/></g>`),
		g.Group(children),
	)
}
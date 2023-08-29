package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerLoadingOneProgressLoadingLoadWaitWaiting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .88v2.5m-5 0l1.84 1.69M1 8.88l2.42-.9m.97 5.14l1.11-2.24m6.5-7.5l-1.84 1.69M13 8.88l-2.42-.9m-.97 5.14L8.5 10.88"/>`),
		g.Group(children),
	)
}
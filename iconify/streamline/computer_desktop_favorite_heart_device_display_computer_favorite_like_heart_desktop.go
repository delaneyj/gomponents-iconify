package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopFavoriteHeartDeviceDisplayComputerFavoriteLikeHeartDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6 11l-1 2.5M8 11l1 2.5m-5 0h6M13 2a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2"/><path d="M7 4c1.17-2.59 3.5-1.29 3.5.65C10.5 7.2 7 8.5 7 8.5S3.5 7.2 3.5 4.6C3.5 2.66 5.83 1.36 7 4Z"/></g>`),
		g.Group(children),
	)
}
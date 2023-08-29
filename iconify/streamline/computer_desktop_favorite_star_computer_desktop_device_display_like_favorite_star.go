package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopFavoriteStarComputerDesktopDeviceDisplayLikeFavoriteStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 2H13a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6"/><path d="M6.45 2.36a.59.59 0 0 1 1.1 0L8.19 4H9.9a.61.61 0 0 1 .56.39a.59.59 0 0 1-.16.61L8.79 6.34l.64 1.28a.58.58 0 0 1-.12.69a.59.59 0 0 1-.7.1L7 7.53l-1.61.88a.59.59 0 0 1-.7-.1a.58.58 0 0 1-.12-.69l.64-1.28L3.7 5a.59.59 0 0 1-.16-.65A.61.61 0 0 1 4.1 4h1.71Z"/></g>`),
		g.Group(children),
	)
}
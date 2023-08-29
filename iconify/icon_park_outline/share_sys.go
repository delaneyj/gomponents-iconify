package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareSys(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10.365 41.632A19.946 19.946 0 0 1 4 27C4 15.954 12.954 7 24 7s20 8.954 20 20a19.945 19.945 0 0 1-6.365 14.632"/><path d="M15.137 36.51A12.965 12.965 0 0 1 11 27c0-7.18 5.82-13 13-13s13 5.82 13 13a12.96 12.96 0 0 1-4.138 9.51"/><path d="M19.91 31.39a6 6 0 1 1 8.18 0"/></g>`),
		g.Group(children),
	)
}
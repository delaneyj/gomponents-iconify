package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceCalendarHeartCalendarDateDayFavoriteHeartLikeMonth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M7 7.5c1-2 3-1 3 .5c0 2-3 3-3 3s-3-1-3-3c0-1.5 2-2.5 3-.5Z"/></g>`),
		g.Group(children),
	)
}
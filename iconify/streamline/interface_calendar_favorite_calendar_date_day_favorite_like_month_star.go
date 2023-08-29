package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceCalendarFavoriteCalendarDateDayFavoriteLikeMonthStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 2.5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-2m-7-2v4m7-4v4m-7-2h5"/><path d="M6.45 5.37a.59.59 0 0 1 1.1 0L8.19 7H9.9a.6.6 0 0 1 .4 1L8.79 9.36l.64 1.28a.58.58 0 0 1-.12.69a.61.61 0 0 1-.7.1L7 10.55l-1.61.88a.61.61 0 0 1-.7-.1a.58.58 0 0 1-.12-.69l.64-1.28L3.7 8a.6.6 0 0 1 .4-1h1.71Z"/></g>`),
		g.Group(children),
	)
}
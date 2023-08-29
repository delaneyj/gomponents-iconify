package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentViewsSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRecentViewsSort0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 5v25h36V5M30 37l-6 6l-6-6m6-7v13"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 21c4.05 0 7.717-2 11-6c-3.283-4-6.95-6-11-6s-7.717 2-11 6c3.283 4 6.95 6 11 6Z"/><path fill="#fff" d="M24 18a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRecentViewsSort0)"/>`),
		g.Group(children),
	)
}
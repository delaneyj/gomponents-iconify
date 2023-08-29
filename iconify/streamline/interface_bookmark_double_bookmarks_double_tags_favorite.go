package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceBookmarkDoubleBookmarksDoubleTagsFavorite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9 13.5l-3.5-3l-3.5 3v-9a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1Z"/><path d="M5 .5h6a1 1 0 0 1 1 1v9"/></g>`),
		g.Group(children),
	)
}
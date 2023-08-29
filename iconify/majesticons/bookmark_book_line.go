package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkBookLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3v9a1 1 0 0 1-1.555.832L11 13.202l-2.445 1.63A1 1 0 0 1 7 14V5H6zm3 0v7.132l1.445-.964a1 1 0 0 1 1.11 0l1.445.964V5H9zM3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6z"/></g>`),
		g.Group(children),
	)
}
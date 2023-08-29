package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11.222L14.667 13l-.89-3.111L16 8l-2.667-.333L12 5l-1.333 2.667L8 8l2.223 1.889L9.333 13z"/><path fill="currentColor" d="M19 21.723V4a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v17.723l7-4.571l7 4.571zM8 8l2.667-.333L12 5l1.333 2.667L16 8l-2.223 1.889l.89 3.111L12 11.222L9.333 13l.89-3.111L8 8z"/>`),
		g.Group(children),
	)
}
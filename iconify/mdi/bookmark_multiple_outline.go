package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMultipleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1h10a2 2 0 0 1 2 2v16l-2-.87V3H7a2 2 0 0 1 2-2m6 19V7H5v13l5-2.18L15 20m0-15a2 2 0 0 1 2 2v16l-7-3l-7 3V7a2 2 0 0 1 2-2h10Z"/>`),
		g.Group(children),
	)
}
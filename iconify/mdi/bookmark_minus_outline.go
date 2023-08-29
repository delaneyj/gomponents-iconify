package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11V9h6v2H9m10-6v16l-7-3l-7 3V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2m-2 0H7v13l5-2.18L17 18V5Z"/>`),
		g.Group(children),
	)
}
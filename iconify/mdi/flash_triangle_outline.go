package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashTriangleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L1 21h22M12 6l7.5 13h-15m9.5-5h-1.5l1.5-3h-4v4h1v3l3-4Z"/>`),
		g.Group(children),
	)
}
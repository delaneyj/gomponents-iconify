package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v4H3V3m1 18V8h16v13H4m10-7v-3h-4v3H7l5 5l5-5h-3Z"/>`),
		g.Group(children),
	)
}
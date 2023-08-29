package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectsVerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v2H2z"/><rect width="6" height="16" x="5" y="6" fill="currentColor" rx="1"/><rect width="6" height="12" x="13" y="6" fill="currentColor" rx="1"/>`),
		g.Group(children),
	)
}
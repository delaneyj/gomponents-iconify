package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 9.83L8.17 12L6 14.17V9.83M4 5v14l7-7m9-7h-2v14h2m-7-9.17L15.17 12L13 14.17V9.83M11 5v14l7-7"/>`),
		g.Group(children),
	)
}
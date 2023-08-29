package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewParallel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3v18h-4V3h4m-5 0v18h-4V3h4M9 3v18H5V3h4Z"/>`),
		g.Group(children),
	)
}
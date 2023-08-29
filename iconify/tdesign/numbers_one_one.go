package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersOneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20V3.5h-1.834L8.1 5.8l1.2 1.6L11 6.125V20h2Z"/>`),
		g.Group(children),
	)
}
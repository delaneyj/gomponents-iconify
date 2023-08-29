package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.364 4.372L7.015 12l13.35 7.628l-.993 1.736L2.984 12l16.388-9.364l.992 1.736Z"/>`),
		g.Group(children),
	)
}
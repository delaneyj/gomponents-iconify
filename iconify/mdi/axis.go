package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.61 21l-1-1.73L11 13.85V3h2v10.85l9.39 5.42l-1 1.73L12 15.58L2.61 21Z"/>`),
		g.Group(children),
	)
}
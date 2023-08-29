package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m5.032 13l.9-3h4.137l.9 3h1.775l-3-10H6.256l-3 10h1.776zm2.4-8h1.137l.9 3H6.532l.9-3z"/>`),
		g.Group(children),
	)
}
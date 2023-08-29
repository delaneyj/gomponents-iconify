package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.62 2c3.35 5.61-8.15 18.15-8.15 18.15L9.6 17.28L4.91 22l-2.14-2.14L20.62 2Z"/>`),
		g.Group(children),
	)
}
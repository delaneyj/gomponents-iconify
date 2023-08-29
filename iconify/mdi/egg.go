package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 14.5A7.5 7.5 0 0 1 12 22a7.5 7.5 0 0 1-7.5-7.5C4.5 10.36 7.86 2 12 2c4.14 0 7.5 8.36 7.5 12.5Z"/>`),
		g.Group(children),
	)
}
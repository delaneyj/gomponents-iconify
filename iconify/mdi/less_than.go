package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.5 4.14l1 1.72L8.97 12l10.53 6.14l-1 1.72L5 12l13.5-7.86Z"/>`),
		g.Group(children),
	)
}
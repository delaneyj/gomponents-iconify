package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.38 2l-5 5l-1.25 8.75L11.38 22l7.5-5l-5-5l3.75-3.75L11.38 2Z"/>`),
		g.Group(children),
	)
}
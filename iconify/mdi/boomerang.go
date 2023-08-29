package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boomerang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2H4c-1.1 0-2 .9-2 2s.9 2 2 2h4l2-4m8 0c2.2 0 4 1.8 4 4v6l-4 2c0-4.4-3.6-8-8-8l2-4h6m0 18v-4l4-2v6c0 1.1-.9 2-2 2s-2-.9-2-2Z"/>`),
		g.Group(children),
	)
}
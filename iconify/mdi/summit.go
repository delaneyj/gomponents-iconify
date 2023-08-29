package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Summit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3h2l5 2l-5 2v3.17L22 21H2l6-8l3.5 4.7l3.5-7.53V3Z"/>`),
		g.Group(children),
	)
}
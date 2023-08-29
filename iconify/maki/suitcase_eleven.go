package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8 3V1.578L7.36 1H3.64L3 1.748V3H1.5l-.5.5v6l.5.5h8l.5-.5v-6L9.5 3H8zM4 2h3v1H4V2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsFlagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M14.303 6l-3-2H6V2H4v20h2v-8h4.697l3 2H20V6z" fill="currentColor"/>`),
		g.Group(children),
	)
}
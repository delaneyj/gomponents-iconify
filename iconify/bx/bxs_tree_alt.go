package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsTreeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18 10l-6-8l-6 8h3l-5 8h7v4h2v-4h7l-5-8h3z" fill="currentColor"/>`),
		g.Group(children),
	)
}
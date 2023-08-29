package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.288 11v2H7.802l3.243 3.243l-1.414 1.414L3.974 12L9.63 6.343l1.414 1.414L7.802 11h14.486ZM3 18V6H1v12h2Z"/>`),
		g.Group(children),
	)
}
package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousChristianEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.5 0v3H2v2h2.5v6h2V5H9V3H6.5V0h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 5.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 10V0H0v10h1ZM5.5 0A2.5 2.5 0 0 0 3 2.5v7.146l2.776 4.361a2.034 2.034 0 0 0 3.536-2.002L8.309 10H12.5A2.5 2.5 0 0 0 15 7.5V4.333L12.5 1a2.5 2.5 0 0 0-2-1h-5Z"/>`),
		g.Group(children),
	)
}
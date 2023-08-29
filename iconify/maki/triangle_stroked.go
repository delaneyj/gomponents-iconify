package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleStroked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.524.5a.77.77 0 0 0-.69.395l-5.5 9.87C1.022 11.307 1.395 12 2 12h11c.605 0 .978-.692.666-1.236L8.166.896A.773.773 0 0 0 7.524.5zM7.5 2.9l4.127 7.47H3.373z"/>`),
		g.Group(children),
	)
}
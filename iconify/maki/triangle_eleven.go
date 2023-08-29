package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.517 1.232a.556.556 0 0 0-.493.268l-4 6.66c-.223.37.044.84.476.84h8c.432 0 .699-.47.476-.84l-4-6.66a.556.556 0 0 0-.459-.268z" fill="currentColor"/>`),
		g.Group(children),
	)
}
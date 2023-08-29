package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.299 13.098a.7.7 0 0 1 .498 1.191l-5.427 5.503a.7.7 0 1 1-.996-.984l4.252-4.31H.703a.7.7 0 0 1 0-1.4H19.3ZM6.619.202a.7.7 0 0 1 .007.99l-4.252 4.31h16.923a.7.7 0 0 1 0 1.4H.7a.7.7 0 0 1-.498-1.191L5.63.208a.7.7 0 0 1 .99-.006Z"/>`),
		g.Group(children),
	)
}
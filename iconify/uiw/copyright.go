package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21Zm.16 3.256c1.08 0 2.117.303 3.003.864a.698.698 0 0 1-.745 1.179a4.211 4.211 0 0 0-2.257-.647c-2.278 0-4.114 1.775-4.114 3.953s1.836 3.953 4.114 3.953c.801 0 1.567-.22 2.224-.627a.698.698 0 0 1 .735 1.187a5.608 5.608 0 0 1-2.96.836c-3.037 0-5.509-2.39-5.509-5.349c0-2.96 2.472-5.349 5.51-5.349Z"/>`),
		g.Group(children),
	)
}
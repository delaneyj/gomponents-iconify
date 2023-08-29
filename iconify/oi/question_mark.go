package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4.47 0c-.85 0-1.48.26-1.88.66c-.4.4-.54.9-.59 1.28l1 .13c.04-.25.12-.5.31-.69C3.5 1.19 3.8 1 4.47 1c.66 0 1.02.16 1.22.34c.2.18.28.4.28.66c0 .83-.34 1.06-.84 1.5c-.5.44-1.16 1.08-1.16 2.25V6h1v-.25c0-.83.31-1.06.81-1.5c.5-.44 1.19-1.08 1.19-2.25c0-.48-.17-1.02-.59-1.41C5.95.2 5.31 0 4.47 0zm-.5 7v1h1V7h-1z"/>`),
		g.Group(children),
	)
}
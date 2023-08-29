package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionArgumentLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10a8 8 0 0 1-14.93 4H.832C2.375 17.532 5.9 20 10 20c5.523 0 10-4.477 10-10S15.523 0 10 0C5.9 0 2.375 2.468.832 6H3.07A8 8 0 0 1 18 10ZM2.062 11H8v4l6-5l-6-5v4H0v2h2.062Z"/>`),
		g.Group(children),
	)
}
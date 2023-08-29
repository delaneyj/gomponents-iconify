package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v8h2V0H0zm7 0v8h1V0H7zM6 2L4 4l2 2V2z"/>`),
		g.Group(children),
	)
}
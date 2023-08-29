package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v8h1V0H0zm6 0v8h2V0H6zM2 2v4l2-2l-2-2z"/>`),
		g.Group(children),
	)
}
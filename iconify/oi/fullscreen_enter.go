package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenEnter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v4l1.5-1.5L3 4l1-1l-1.5-1.5L4 0H0zm5 4L4 5l1.5 1.5L4 8h4V4L6.5 5.5L5 4z"/>`),
		g.Group(children),
	)
}
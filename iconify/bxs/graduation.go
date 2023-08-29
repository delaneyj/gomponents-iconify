package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Graduation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7v1l11 4l9-4V7L11 4z"/><path fill="currentColor" d="M4 11v4.267c0 1.621 4.001 3.893 9 3.734c4-.126 6.586-1.972 7-3.467c.024-.089.037-.178.037-.268V11L13 14l-5-1.667v3.213l-1-.364V12l-3-1z"/>`),
		g.Group(children),
	)
}
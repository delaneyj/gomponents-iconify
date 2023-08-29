package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h12v2H6zm5-14v8.586L6.707 8.293L5.293 9.707L12 16.414l6.707-6.707l-1.414-1.414L13 12.586V4z"/>`),
		g.Group(children),
	)
}
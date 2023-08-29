package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLastLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 1h2v18h-2zM3.5 2.5L11 10l-7.5 7.5L5 19l9-9l-9-9z"/>`),
		g.Group(children),
	)
}
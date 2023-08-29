package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.47 13.77l-1.41-1.42l5.66-5.65l-1.42-1.42l-5.65 5.66l-1.42-1.41l5.66-5.66l-1.42-1.42l-6.36 6.37a3 3 0 0 0 0 4.24l.71.71l-6.37 6.36l1.42 1.42l6.36-6.37l.71.71a3 3 0 0 0 4.24 0l6.37-6.36l-1.42-1.42z"/>`),
		g.Group(children),
	)
}
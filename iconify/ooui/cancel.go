package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0zM2 10a8 8 0 0 1 1.69-4.9L14.9 16.31A8 8 0 0 1 2 10zm14.31 4.9L5.1 3.69A8 8 0 0 1 16.31 14.9z"/>`),
		g.Group(children),
	)
}
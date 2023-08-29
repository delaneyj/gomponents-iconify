package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M6 0L5 1l2 2l1-1l-2-2zM4 2L0 6v2h2l4-4l-2-2z"/>`),
		g.Group(children),
	)
}
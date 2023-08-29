package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.14 2.27a1 1 0 0 0-1.41 0l-10 10a1 1 0 0 0 0 1.41L4 14l-3 4h5l1-1l.29.29a1 1 0 0 0 1.41 0l10-10a1 1 0 0 0 .03-1.43zM7 15l-2-2l9-9l2 2z"/>`),
		g.Group(children),
	)
}
package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 3H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"/><path fill="currentColor" d="M21 19v-9a2 2 0 0 0-2-2h-1v8a2 2 0 0 1-2 2H8v1a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2z"/>`),
		g.Group(children),
	)
}
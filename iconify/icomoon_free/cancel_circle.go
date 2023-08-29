package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13z"/><path fill="currentColor" d="M10.5 4L8 6.5L5.5 4L4 5.5L6.5 8L4 10.5L5.5 12L8 9.5l2.5 2.5l1.5-1.5L9.5 8L12 5.5z"/>`),
		g.Group(children),
	)
}
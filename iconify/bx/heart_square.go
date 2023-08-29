package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.998 17l4.186-4.186a2.745 2.745 0 0 0 0-3.907a2.746 2.746 0 0 0-3.907 0l-.278.279l-.279-.279a2.746 2.746 0 0 0-3.907 0a2.746 2.746 0 0 0 0 3.907L11.998 17z"/><path fill="currentColor" d="M21 4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4zm-2 15H5V5h14v14z"/>`),
		g.Group(children),
	)
}
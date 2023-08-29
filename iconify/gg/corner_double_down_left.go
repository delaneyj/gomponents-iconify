package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.295 7.694l4.84-5.056l5.057 4.84l-1.383 1.445l-2.462-2.357l.162 6.034a4.8 4.8 0 0 1-4.67 4.927l-5.925.16l2.294 2.246l-1.4 1.43l-5-4.9l4.898-5l1.429 1.4l-2.377 2.427l6.017-.162a2.4 2.4 0 0 0 2.335-2.463l-.158-5.898l-2.212 2.31l-1.445-1.383Z"/>`),
		g.Group(children),
	)
}
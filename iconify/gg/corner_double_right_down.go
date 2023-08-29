package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.694 12.705l-5.056-4.84l4.84-5.057L8.923 4.19L6.566 6.653L12.6 6.49a4.8 4.8 0 0 1 4.927 4.669l.16 5.926l2.246-2.294l1.43 1.4l-4.9 5l-5-4.898l1.4-1.429l2.427 2.378l-.162-6.018a2.4 2.4 0 0 0-2.463-2.335l-5.898.158l2.31 2.212l-1.383 1.445Z"/>`),
		g.Group(children),
	)
}
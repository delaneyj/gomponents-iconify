package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.27 1.27L16.11 18H13.9l-3.7 3.71c-.2.19-.45.29-.7.29H9c-.55 0-1-.45-1-1v-3H4a2 2 0 0 1-2-2V3.9L1.11 3l1.28-1.27l19.72 19.73M22 16V4a2 2 0 0 0-2-2H5.2l15.75 15.75c.62-.34 1.05-.99 1.05-1.75Z"/>`),
		g.Group(children),
	)
}
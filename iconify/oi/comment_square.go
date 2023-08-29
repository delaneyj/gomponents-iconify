package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V5.9c0 .05.04.09.09.09H6l2 2V.08c0-.06-.04-.09-.09-.09H.1z"/>`),
		g.Group(children),
	)
}
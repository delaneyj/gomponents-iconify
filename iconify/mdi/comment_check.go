package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h4v3c0 .55.45 1 1 1h.5c.25 0 .5-.1.7-.29L13.9 18H20c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2m-9.53 12L7 10.5l1.4-1.41l2.07 2.08L15.6 6L17 7.41L10.47 14Z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-4-9A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20Zm0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20Z"/>`),
		g.Group(children),
	)
}
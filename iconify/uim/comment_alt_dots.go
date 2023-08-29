package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v10a3.003 3.003 0 0 0 3 3h11.586l3.707 3.707A1 1 0 0 0 22 21V5a3.003 3.003 0 0 0-3-3ZM8 11a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z" opacity=".5"/><circle cx="8" cy="10" r="1" fill="currentColor"/><circle cx="12" cy="10" r="1" fill="currentColor"/><circle cx="16" cy="10" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}
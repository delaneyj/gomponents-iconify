package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 13H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M12 2a10 10 0 0 0-7.743 16.33l-1.964 1.963A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20ZM9 7h6a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm6 10H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Zm2-4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z" opacity=".5"/><path fill="currentColor" d="M15 17H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2zm0-8H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}
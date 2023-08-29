package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10.017 10.017 0 0 0-7 2.877V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2H6.218A7.992 7.992 0 1 1 12 20a1 1 0 0 0 0 2a10 10 0 0 0 0-20Z"/><path fill="currentColor" d="M14 13h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M12 4a8.008 8.008 0 0 0-5.782 2.5H8.5a1 1 0 0 1 0 2H4a.989.989 0 0 1-.976-.88A9.977 9.977 0 0 0 12 22a1 1 0 0 1 0-2a8 8 0 0 0 0-16Zm2 9h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z" opacity=".5"/>`),
		g.Group(children),
	)
}
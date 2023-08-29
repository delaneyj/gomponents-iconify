package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h4.56a1 1 0 0 1 .95.68l.54 1.64A1 1 0 0 0 11 7h7a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-6.28l-.32-1a3 3 0 0 0-2.84-2H4a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1Zm17 11.18V14a3 3 0 0 0-6 0v1.18A3 3 0 0 0 13 18v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-2-2.82ZM17 14a1 1 0 0 1 2 0v1h-2Zm4 6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}
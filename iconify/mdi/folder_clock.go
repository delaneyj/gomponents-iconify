package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h7.26c1.31 1.88 3.45 3 5.74 3a7 7 0 0 0 7-7c0-1.83-.72-3.58-2-4.89V8a2 2 0 0 0-2-2h-8L9 4H3m13 7a5 5 0 0 1 5 5a5 5 0 0 1-5 5a5 5 0 0 1-5-5a5 5 0 0 1 5-5m-1 1v5l3.61 2.16l.75-1.22l-2.86-1.69V12H15Z"/>`),
		g.Group(children),
	)
}
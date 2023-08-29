package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderTableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-8l-2-2H4m0 4h16v10H4V8m8 1v2h3V9h-3m4 0v2h3V9h-3m-4 3v2h3v-2h-3m4 0v2h3v-2h-3m-4 3v2h3v-2h-3m4 0v2h3v-2h-3Z"/>`),
		g.Group(children),
	)
}
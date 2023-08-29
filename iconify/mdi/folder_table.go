package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-8l-2-2H4m8 5h3v2h-3V9m4 0h3v2h-3V9m-4 3h3v2h-3v-2m4 0h3v2h-3v-2m-4 3h3v2h-3v-2m4 0h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}
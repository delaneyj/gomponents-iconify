package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18H4V8h16v10M12 6l-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16c1.11 0 2-.89 2-2V8a2 2 0 0 0-2-2h-8m-1 8v-2h4V9l4 4l-4 4v-3h-4Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20h16c1.11 0 2-.89 2-2V6c0-1.11-.89-2-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2m0-9h16v7H4v-7Z"/>`),
		g.Group(children),
	)
}
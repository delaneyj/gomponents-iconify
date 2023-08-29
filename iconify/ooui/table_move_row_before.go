package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMoveRowBefore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 9H6l4-5l4 5h-3v5H9z"/><path fill="currentColor" d="M2 0h16v20H2zm2 2v8h4v5h4v-5h4V2z"/>`),
		g.Group(children),
	)
}
package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMoveRowAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 16l-4-5h3V6h2v5h3z"/><path fill="currentColor" d="M2 0v20h16V0zm2 10h4V5h4v5h4v8H4z"/>`),
		g.Group(children),
	)
}
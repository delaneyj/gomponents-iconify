package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19 10l-6-5v4H6v2h7v4zM6 2V1H4.5a1.49 1.49 0 0 0-1 .39a1.49 1.49 0 0 0-1-.39H1v1h1.5a.5.5 0 0 1 .5.5v15a.5.5 0 0 1-.5.5H1v1h1.5a1.49 1.49 0 0 0 1-.39a1.49 1.49 0 0 0 1 .39H6v-1H4.5a.5.5 0 0 1-.5-.5v-15a.5.5 0 0 1 .5-.5z"/>`),
		g.Group(children),
	)
}
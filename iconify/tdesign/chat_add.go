package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 0v3h3v2h-3v3h-2V5h-3V3h3V0h2ZM1.5 2H14v2H3.5v14.296L6.124 16H20.5v-6h2v8H6.876L1.5 22.704V2Z"/>`),
		g.Group(children),
	)
}
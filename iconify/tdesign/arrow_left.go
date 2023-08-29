package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 13H7.914l4.5 4.5L11 18.914L4.086 12L11 5.086L12.414 6.5l-4.5 4.5H19.5v2Z"/>`),
		g.Group(children),
	)
}
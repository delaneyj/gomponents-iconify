package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 2.586v18.828L6.086 12L15.5 2.586ZM8.914 12l4.586 4.586V7.414L8.914 12Z"/>`),
		g.Group(children),
	)
}
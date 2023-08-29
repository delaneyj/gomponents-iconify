package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0ZM12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1Zm5.5 12H9.914l3 3l-1.414 1.414L6.086 12L11.5 6.586L12.914 8l-3 3H17.5v2Z"/>`),
		g.Group(children),
	)
}
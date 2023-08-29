package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Zm9-11C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1ZM6.5 13h7.586l-3 3l1.414 1.414L17.914 12L12.5 6.586L11.086 8l3 3H6.5v2Z"/>`),
		g.Group(children),
	)
}
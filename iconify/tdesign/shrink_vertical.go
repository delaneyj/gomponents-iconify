package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShrinkVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1.5v5.086l2.5-2.5L16.914 5.5L12 10.414L7.086 5.5L8.5 4.086l2.5 2.5V1.5h2ZM21 13H3v-2h18v2Zm-9 .586l4.914 4.914l-1.414 1.414l-2.5-2.5V22.5h-2v-5.086l-2.5 2.5L7.086 18.5L12 13.586Z"/>`),
		g.Group(children),
	)
}
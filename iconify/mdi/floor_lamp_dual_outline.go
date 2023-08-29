package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampDualOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.73 10.06l-1.56 1.18l.83.96V20H8v2h8v-2h-3v-7.7l.85-1.04l-1.58-1.19l-.54-.01M8.32 5.31l.91 2.88l-.23.18L6.47 6.7l1.85-1.39M9.39 2L3 6.81l6.08 4.01l2.47-1.89L9.39 2m6.28 3.3l1.86 1.4L15 8.37l-.23-.17l.9-2.9M14.61 2l-2.17 6.95l2.48 1.87L21 6.81L14.61 2Z"/>`),
		g.Group(children),
	)
}
package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func University(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M24.999 27.381c-5.406 0-9.999 1.572-12.999 4.036V36h26v-4.583c-3-2.464-7.594-4.036-13.001-4.036zm23.871-2.352L24.936 14L1.012 25.029L5 26.854v2.807c-1 .207-1.003.731-1.003 1.354c0 .368.122.799.354 1.057L2.983 35h4.88l-1.356-2.93c.228-.258.415-.638.415-1.006c0-.622-.922-1.197-.922-1.404v-2.337l5 2.246v-.199c3-2.609 8.271-4.265 13.998-4.265C30.727 25.105 36 26.761 39 29.37v.199l9.87-4.54z"/>`),
		g.Group(children),
	)
}
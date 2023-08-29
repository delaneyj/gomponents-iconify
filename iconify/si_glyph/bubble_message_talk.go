package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.939 1C2.666 1 .009 2.987.009 5.438c0 2.236 2.215 4.082 5.092 4.387L3.88 12.26l4.249-2.7C10.318 8.906 12 7.309 12 5.438C12 2.988 9.213 1 5.939 1zm10.008 8.89c0-1.124-1.062-2.288-2.289-2.868c-.344 1.95-1.924 3.745-4.417 4.447l-1.187.642c.454.34 1.01.611 1.634.788l3.638 1.971l-1.303-1.776c2.217-.225 3.924-1.571 3.924-3.204z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallBellFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M16 176a8 8 0 0 1 8-8h8v-16a96.12 96.12 0 0 1 88-95.66V40h-16a8 8 0 0 1 0-16h48a8 8 0 0 1 0 16h-16v16.34A96.12 96.12 0 0 1 224 152v16h8a8 8 0 0 1 0 16H24a8 8 0 0 1-8-8Zm216 24H24a8 8 0 0 0 0 16h208a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}
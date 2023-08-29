package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSlightRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20q-.425 0-.713-.288T9 19v-6.575q0-.4.15-.763t.425-.637L14.6 6h-1.25q-.425 0-.712-.287T12.35 5q0-.425.288-.713T13.35 4H17q.425 0 .713.288T18 5v3.65q0 .425-.288.713T17 9.65q-.425 0-.713-.288T16 8.65V7.4l-5 5V19q0 .425-.288.713T10 20Z"/>`),
		g.Group(children),
	)
}
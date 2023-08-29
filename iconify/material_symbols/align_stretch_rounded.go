package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStretchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10q-.425 0-.713-.288T7 9V4H3q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4h-4v5q0 .425-.288.713T16 10H8ZM3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h4v-5q0-.425.288-.713T8 14h8q.425 0 .713.288T17 15v5h4q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatParagraphRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20q-.425 0-.713-.288T9 19v-5q-2.075 0-3.538-1.463T4 9q0-2.075 1.463-3.538T9 4h8q.425 0 .713.288T18 5q0 .425-.288.713T17 6h-1v13q0 .425-.288.713T15 20q-.425 0-.713-.288T14 19V6h-3v13q0 .425-.288.713T10 20Z"/>`),
		g.Group(children),
	)
}
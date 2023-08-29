package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextClipRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.425 0-.713-.288T4 19V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v14q0 .425-.288.713T5 20Zm14 0q-.425 0-.713-.288T18 19v-6H9q-.425 0-.713-.288T8 12q0-.425.288-.713T9 11h9V5q0-.425.288-.713T19 4q.425 0 .713.288T20 5v14q0 .425-.288.713T19 20Z"/>`),
		g.Group(children),
	)
}
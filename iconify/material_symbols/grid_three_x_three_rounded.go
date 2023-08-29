package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridThreeXThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16H5q-.425 0-.713-.288T4 15q0-.425.288-.713T5 14h3v-4H5q-.425 0-.713-.288T4 9q0-.425.288-.713T5 8h3V5q0-.425.288-.713T9 4q.425 0 .713.288T10 5v3h4V5q0-.425.288-.713T15 4q.425 0 .713.288T16 5v3h3q.425 0 .713.288T20 9q0 .425-.288.713T19 10h-3v4h3q.425 0 .713.288T20 15q0 .425-.288.713T19 16h-3v3q0 .425-.288.713T15 20q-.425 0-.713-.288T14 19v-3h-4v3q0 .425-.288.713T9 20q-.425 0-.713-.288T8 19v-3Zm2-2h4v-4h-4v4Z"/>`),
		g.Group(children),
	)
}
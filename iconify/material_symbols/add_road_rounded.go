package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddRoadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12V5q0-.425.288-.713T19 4q.425 0 .713.288T20 5v7q0 .425-.288.713T19 13q-.425 0-.713-.288T18 12ZM4 19V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v14q0 .425-.288.713T5 20q-.425 0-.713-.288T4 19Zm7-12V5q0-.425.288-.713T12 4q.425 0 .713.288T13 5v2q0 .425-.288.713T12 8q-.425 0-.713-.288T11 7Zm0 6v-2q0-.425.288-.713T12 10q.425 0 .713.288T13 11v2q0 .425-.288.713T12 14q-.425 0-.713-.288T11 13Zm0 6v-2q0-.425.288-.713T12 16q.425 0 .713.288T13 17v2q0 .425-.288.713T12 20q-.425 0-.713-.288T11 19Zm7 1h-2q-.425 0-.713-.288T15 19q0-.425.288-.713T16 18h2v-2q0-.425.288-.713T19 15q.425 0 .713.288T20 16v2h2q.425 0 .713.288T23 19q0 .425-.288.713T22 20h-2v2q0 .425-.288.713T19 23q-.425 0-.713-.288T18 22v-2Z"/>`),
		g.Group(children),
	)
}
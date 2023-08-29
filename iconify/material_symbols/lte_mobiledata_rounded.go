package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LteMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15V9q0-.425.288-.713T5 8q.425 0 .713.288T6 9v5h2q.425 0 .713.288T9 15q0 .425-.288.713T8 16H5q-.425 0-.713-.288T4 15Zm7-5h-1q-.425 0-.713-.288T9 9q0-.425.288-.713T10 8h4q.425 0 .713.288T15 9q0 .425-.288.713T14 10h-1v5q0 .425-.288.713T12 16q-.425 0-.713-.288T11 15v-5Zm5 5V9q0-.425.288-.713T17 8h3q.425 0 .713.288T21 9q0 .425-.288.713T20 10h-2v1h2q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-2v1h2q.425 0 .713.288T21 15q0 .425-.288.713T20 16h-3q-.425 0-.713-.288T16 15Z"/>`),
		g.Group(children),
	)
}
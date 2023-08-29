package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.425 0-.713-.288T3 18q0-.425.288-.713T4 17h4q.425 0 .713.288T9 18q0 .425-.288.713T8 19H4ZM4 7q-.425 0-.713-.288T3 6q0-.425.288-.713T4 5h8q.425 0 .713.288T13 6q0 .425-.288.713T12 7H4Zm8 14q-.425 0-.713-.288T11 20v-4q0-.425.288-.713T12 15q.425 0 .713.288T13 16v1h7q.425 0 .713.288T21 18q0 .425-.288.713T20 19h-7v1q0 .425-.288.713T12 21Zm-4-6q-.425 0-.713-.288T7 14v-1H4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h3v-1q0-.425.288-.713T8 9q.425 0 .713.288T9 10v4q0 .425-.288.713T8 15Zm4-2q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11h8q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-8Zm4-4q-.425 0-.713-.288T15 8V4q0-.425.288-.713T16 3q.425 0 .713.288T17 4v1h3q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-3v1q0 .425-.288.713T16 9Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublishRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 11.85l-1.875 1.875q-.3.3-.713.288T7.7 13.7q-.275-.3-.288-.7t.288-.7l3.6-3.6q.15-.15.325-.212T12 8.425q.2 0 .375.063t.325.212l3.6 3.6q.3.3.288.7t-.288.7q-.3.3-.713.313t-.712-.288L13 11.85V19q0 .425-.288.713T12 20q-.425 0-.713-.288T11 19v-7.15ZM4 8V6q0-.825.588-1.413T6 4h12q.825 0 1.413.588T20 6v2q0 .425-.288.713T19 9q-.425 0-.713-.288T18 8V6H6v2q0 .425-.288.713T5 9q-.425 0-.713-.288T4 8Z"/>`),
		g.Group(children),
	)
}
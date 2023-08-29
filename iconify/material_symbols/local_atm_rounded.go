package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalAtmRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14h-3q-.425 0-.713.288T9 15q0 .425.288.713T10 16h1q0 .425.288.713T12 17q.4 0 .563-.363T13 16h1q.425 0 .713-.288T15 15v-3q0-.425-.288-.713T14 11h-3v-1h3q.425 0 .713-.288T15 9q0-.425-.288-.713T14 8h-1q0-.425-.288-.713T12 7q-.4 0-.563.363T11 8h-1q-.425 0-.713.288T9 9v3q0 .425.288.713T10 13h3v1Zm-9 6q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}
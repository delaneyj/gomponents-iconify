package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfirmationNumber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm0-4q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm0-4q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Zm8 11H4q-.825 0-1.413-.588T2 18v-4q.825 0 1.413-.588T4 12q0-.825-.588-1.413T2 10V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4q-.825 0-1.413.588T20 12q0 .825.588 1.413T22 14v4q0 .825-.588 1.413T20 20Z"/>`),
		g.Group(children),
	)
}
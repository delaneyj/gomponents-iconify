package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfirmationNumberOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm0-4q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm0-4q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9ZM2 20v-6q.825 0 1.413-.588T4 12q0-.825-.588-1.413T2 10V4h20v6q-.825 0-1.413.588T20 12q0 .825.588 1.413T22 14v6H2Zm2-2h16v-2.55q-.925-.55-1.463-1.462T18 12q0-1.075.537-1.988T20 8.55V6H4v2.55q.925.55 1.463 1.463T6 12q0 1.075-.537 1.988T4 15.45V18Zm8-6Z"/>`),
		g.Group(children),
	)
}
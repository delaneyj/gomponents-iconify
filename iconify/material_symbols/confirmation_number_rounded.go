package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfirmationNumberRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-3.375q0-.275.175-.475t.45-.25q.6-.2.988-.725T4 12q0-.65-.388-1.175t-.987-.725q-.275-.05-.45-.25T2 9.375V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v3.375q0 .275-.175.475t-.45.25q-.6.2-.988.725T20 12q0 .65.388 1.175t.987.725q.275.05.45.25t.175.475V18q0 .825-.588 1.413T20 20H4Zm8-3q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm0-4q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm0-4q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Z"/>`),
		g.Group(children),
	)
}
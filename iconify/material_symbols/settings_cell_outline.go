package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsCellOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 24q-.425 0-.713-.288T7 23q0-.425.288-.713T8 22q.425 0 .713.288T9 23q0 .425-.288.713T8 24Zm4 0q-.425 0-.713-.288T11 23q0-.425.288-.713T12 22q.425 0 .713.288T13 23q0 .425-.288.713T12 24Zm4 0q-.425 0-.713-.288T15 23q0-.425.288-.713T16 22q.425 0 .713.288T17 23q0 .425-.288.713T16 24Zm-8-4q-.825 0-1.413-.588T6 18V2q0-.825.588-1.413T8 0h8q.825 0 1.413.588T18 2v16q0 .825-.588 1.413T16 20H8Zm0-2h8v-1H8v1Zm0-3h8V5H8v10ZM8 3h8V2H8v1Zm0 0V2v1Zm0 15v-1v1Z"/>`),
		g.Group(children),
	)
}
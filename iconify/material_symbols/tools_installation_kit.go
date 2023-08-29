package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsInstallationKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20q-3.35 0-5.675-2.325T1 12q0-3.35 2.325-5.675T9 4h6q3.35 0 5.675 2.325T23 12q0 3.35-2.325 5.675T15 20H9Zm3-6q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Zm-1-5h2q.425 0 .713-.288T14 8q0-.425-.288-.713T13 7h-2q-.425 0-.713.288T10 8q0 .425.288.713T11 9Zm0 8h2q.425 0 .713-.288T14 16q0-.425-.288-.713T13 15h-2q-.425 0-.713.288T10 16q0 .425.288.713T11 17Zm5-3q.425 0 .713-.288T17 13v-2q0-.425-.288-.713T16 10q-.425 0-.713.288T15 11v2q0 .425.288.713T16 14Zm-8 0q.425 0 .713-.288T9 13v-2q0-.425-.288-.713T8 10q-.425 0-.713.288T7 11v2q0 .425.288.713T8 14Z"/>`),
		g.Group(children),
	)
}
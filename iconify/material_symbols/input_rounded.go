package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-2q0-.425.288-.713T3 15q.425 0 .713.288T4 16v2h16V6H4v2q0 .425-.288.713T3 9q-.425 0-.713-.288T2 8V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm8.175-7H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h9.175L10.8 9.65q-.3-.275-.288-.688t.288-.712q.3-.3.713-.313t.712.288L15.3 11.3q.3.3.3.7t-.3.7l-3.075 3.075q-.3.3-.712.3t-.713-.3q-.275-.3-.288-.713t.288-.712L12.175 13Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Problem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16q.425 0 .713-.288T8 15q0-.425-.288-.713T7 14q-.425 0-.713.288T6 15q0 .425.288.713T7 16Zm-1-3h2V8H6v5Zm4 2h8v-2h-8v2Zm0-4h8V9h-8v2Zm-6 9q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}
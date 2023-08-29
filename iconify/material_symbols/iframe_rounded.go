package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IframeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15q-.425 0-.713-.288T10 14v-3q0-.425.288-.713T11 10h6q.425 0 .713.288T18 11v3q0 .425-.288.713T17 15h-6Zm-7 5q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V8H4v10Z"/>`),
		g.Group(children),
	)
}
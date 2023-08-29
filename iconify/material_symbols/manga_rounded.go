package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MangaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm4.1-2H20V7.025L17 8l-2.4-.775q-.3-.1-.6-.013t-.5.363L12 9.625l-2.4.775q-.325.1-.5.363t-.175.587v2.525L7.45 15.9q-.2.275-.2.588t.2.587L8.1 18Z"/>`),
		g.Group(children),
	)
}
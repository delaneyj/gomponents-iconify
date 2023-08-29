package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 15q-.425 0-.713-.288T14.5 14v-4q0-.425.288-.713T15.5 9H18q.425 0 .713.288T19 10v4q0 .425-.288.713T18 15h-2.5Zm.5-1.5h1.5v-3H16v3ZM5 15v-2.5q0-.425.288-.713T6 11.5h2v-1H5V9h3.5q.425 0 .713.288T9.5 10v1.5q0 .425-.288.713T8.5 12.5h-2v1h3V15H5Zm6.25-4V9.5h1.5V11h-1.5Zm0 3.5V13h1.5v1.5h-1.5ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h3V2h2v2h6V2h2v2h3q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h7.25v-1.5h1.5V18H20V6h-7.25v1.5h-1.5V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
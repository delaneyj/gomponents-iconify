package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorbellThreePOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22V2h14v20H5Zm2-2h10V4H7v16Zm5-1q.825 0 1.413-.588T14 17q0-.825-.588-1.413T12 15q-.825 0-1.413.588T10 17q0 .825.588 1.413T12 19Zm0-1q-.425 0-.713-.288T11 17q0-.425.288-.713T12 16q.425 0 .713.288T13 17q0 .425-.288.713T12 18Zm0-4.5q.45 0 .725-.275T13 12.5h-2q0 .45.275.725T12 13.5ZM8 12h8v-1h-1V8.7q0-1.125-.575-2.013T12.8 5.5V4.4h-1.6v1.1q-1.05.375-1.625 1.225T9 8.7V11H8v1Zm-1 8V4v16Z"/>`),
		g.Group(children),
	)
}
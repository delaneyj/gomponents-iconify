package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopicOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm0-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm3-6h10q.425 0 .713-.288T18 11q0-.425-.288-.713T17 10H7q-.425 0-.713.288T6 11q0 .425.288.713T7 12Zm0 4h6q.425 0 .713-.288T14 15q0-.425-.288-.713T13 14H7q-.425 0-.713.288T6 15q0 .425.288.713T7 16Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestPageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17v.5q0 .2.15.35t.35.15h1q.2 0 .35-.15t.15-.35V17h1q.425 0 .713-.288T15 16v-3q0-.425-.288-.713T14 12h-3v-1h3q.425 0 .713-.288T15 10q0-.425-.288-.713T14 9h-1v-.5q0-.2-.15-.35T12.5 8h-1q-.2 0-.35.15T11 8.5V9h-1q-.425 0-.713.288T9 10v3q0 .425.288.713T10 14h3v1h-3q-.425 0-.713.288T9 16q0 .425.288.713T10 17h1Zm-5 5q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm0-2h12V8.85L13.15 4H6v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}
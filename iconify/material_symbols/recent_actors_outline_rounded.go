package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentActorsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h10q.825 0 1.413.588T15 7v10q0 .825-.588 1.413T13 19H3Zm0-3.65q1.1-.65 2.35-1T8 14q1.4 0 2.65.35t2.35 1V7H3v8.35ZM8 16q-1.025 0-2 .25T4.15 17h7.7q-.875-.5-1.85-.75T8 16Zm0-2.75q-1.125 0-1.938-.813T5.25 10.5q0-1.125.813-1.938T8 7.75q1.125 0 1.938.813t.812 1.937q0 1.125-.813 1.938T8 13.25Zm0-1.85q.375 0 .638-.263T8.9 10.5q0-.375-.263-.638T8 9.6q-.375 0-.638.263T7.1 10.5q0 .375.263.638T8 11.4ZM18 19q-.425 0-.713-.288T17 18V6q0-.425.288-.713T18 5q.425 0 .713.288T19 6v12q0 .425-.288.713T18 19Zm4 0q-.425 0-.713-.288T21 18V6q0-.425.288-.713T22 5q.425 0 .713.288T23 6v12q0 .425-.288.713T22 19ZM8 10.5ZM8 17Z"/>`),
		g.Group(children),
	)
}
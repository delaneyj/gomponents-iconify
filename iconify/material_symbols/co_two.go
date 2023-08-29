package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15q-.425 0-.713-.288T10 14v-4q0-.425.288-.713T11 9h3q.425 0 .713.288T15 10v4q0 .425-.288.713T14 15h-3Zm.5-1.5h2v-3h-2v3ZM4 15q-.425 0-.713-.288T3 14v-4q0-.425.288-.713T4 9h3q.425 0 .713.288T8 10v1H6.5v-.5h-2v3h2V13H8v1q0 .425-.288.713T7 15H4Zm13 3v-2.5q0-.425.288-.713T18 14.5h2v-1h-3V12h3.5q.425 0 .713.288T21.5 13v1.5q0 .425-.288.713t-.712.287h-2v1h3V18H17Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-3H7q-.825 0-1.413-.588T5 17V7H2q-.425 0-.713-.288T1 6q0-.425.288-.713T2 5h3V2q0-.425.288-.713T6 1q.425 0 .713.288T7 2v15h15q.425 0 .713.288T23 18q0 .425-.288.713T22 19h-3v3q0 .425-.288.713T18 23q-.425 0-.713-.288T17 22Zm0-7V7H9V5h8q.825 0 1.413.588T19 7v8h-2Z"/>`),
		g.Group(children),
	)
}
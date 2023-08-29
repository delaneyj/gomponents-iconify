package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.425 14.675H12q.425 0 .713-.288t.287-.712V11.35h1.575q.425 0 .713-.287t.287-.713V8H17q.425 0 .713-.287T18 7q0-.425-.288-.713T17 6h-2.425q-.425 0-.713.288T13.575 7v2.325H12q-.425 0-.713.288t-.287.712v2.325H9.425q-.425 0-.713.288t-.287.712V16H7q-.425 0-.712.288T6 17q0 .425.288.713T7 18h2.425q.425 0 .713-.288t.287-.712v-2.325ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
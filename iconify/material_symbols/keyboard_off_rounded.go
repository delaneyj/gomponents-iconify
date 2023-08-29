package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.15 21.8L2.175 4.85q-.3-.3-.287-.712t.312-.713q.3-.3.713-.3t.712.3L20.575 20.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM16.175 16l-2-2H9q-.425 0-.713.288T8 15q0 .425.288.713T9 16h7.175ZM5 13h2v-2H5v2Zm3 0h2v-2H8v2Zm9 0h2v-2h-2v2ZM5 10h2V8H5v2Zm9 0h2V8h-2v2Zm3 0h2V8h-2v2ZM4 19q-.825 0-1.413-.588T2 17V7q0-.825.588-1.413T4 5h1.175l14 14H4Zm17.4-.575l-5.4-5.45V11h-2l-1-1.025V8h-1.975l-3-3H20q.825 0 1.413.588T22 7v10.025q0 .425-.163.775t-.437.625Z"/>`),
		g.Group(children),
	)
}
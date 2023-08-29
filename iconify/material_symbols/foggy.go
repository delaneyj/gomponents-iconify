package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foggy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 19q-.425 0-.713-.288T17 18q0-.425.288-.713T18 17q.425 0 .713.288T19 18q0 .425-.288.713T18 19ZM7 22q-.425 0-.713-.288T6 21q0-.425.288-.713T7 20q.425 0 .713.288T8 21q0 .425-.288.713T7 22Zm-1-3q-.425 0-.713-.288T5 18q0-.425.288-.713T6 17h9q.425 0 .713.288T16 18q0 .425-.288.713T15 19H6Zm4 3q-.425 0-.713-.288T9 21q0-.425.288-.713T10 20h7q.425 0 .713.288T18 21q0 .425-.288.713T17 22h-7Zm-2.5-6q-2.275 0-3.888-1.613T2 10.5q0-2.075 1.375-3.625t3.4-1.825q.8-1.425 2.188-2.238T12 2q2.25 0 3.913 1.438t2.012 3.587q1.725.15 2.9 1.425T22 11.5q0 1.875-1.312 3.188T17.5 16h-10Z"/>`),
		g.Group(children),
	)
}
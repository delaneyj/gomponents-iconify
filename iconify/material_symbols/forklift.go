package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forklift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-1.25 0-2.125-.875T1 18q0-.65.25-1.238T2 15.75V11h2V5h8l4.7 11.075q.15.35.225.7T17 17.5q0 1.45-1.025 2.475T13.5 21q-1.025 0-1.888-.537T10.326 19h-3.5q-.325.9-1.1 1.45T4 21Zm14-1V4h2v14h3v2h-5ZM4 19q.425 0 .713-.288T5 18q0-.425-.288-.713T4 17q-.425 0-.713.288T3 18q0 .425.288.713T4 19Zm9.5 0q.625 0 1.063-.438T15 17.5q0-.625-.438-1.063T13.5 16q-.625 0-1.063.438T12 17.5q0 .625.438 1.063T13.5 19Zm-4.575-5h4.725l-2.975-7H6v4l2.925 3Z"/>`),
		g.Group(children),
	)
}
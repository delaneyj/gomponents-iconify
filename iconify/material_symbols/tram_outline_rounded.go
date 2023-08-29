package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TramOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.5V8q0-2.425 2.125-3.175T11 4l.75-1.5h-4q-.325 0-.537-.213T7 1.75q0-.325.213-.537T7.75 1h8.5q.325 0 .537.213T17 1.75q0 .325-.213.537t-.537.213h-2.5L13 4q2.975.075 4.988.813T20 8v9.5q0 1.475-1.012 2.488T16.5 21l.5.5q.425.425.2.963t-.825.537q-.175 0-.338-.063t-.287-.187L14 21h-4l-1.75 1.75q-.125.125-.288.188T7.626 23q-.575 0-.813-.537T7 21.5l.5-.5q-1.475 0-2.488-1.012T4 17.5ZM16.5 14H6h12h-1.5ZM12 18q.625 0 1.063-.438T13.5 16.5q0-.625-.438-1.063T12 15q-.625 0-1.063.438T10.5 16.5q0 .625.438 1.063T12 18Zm-.05-11h5.7H6.4h5.55ZM6 12h12V9H6v3Zm1.5 7h9q.65 0 1.075-.425T18 17.5V14H6v3.5q0 .65.425 1.075T7.5 19Zm4.45-13q-3.35 0-4.3.363T6.4 7h11.25q-.3-.35-1.3-.675T11.95 6Z"/>`),
		g.Group(children),
	)
}
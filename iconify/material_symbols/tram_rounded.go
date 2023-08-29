package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TramRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.5V8q0-2.425 2.125-3.175T11 4l.75-1.5h-4q-.325 0-.537-.213T7 1.75q0-.325.213-.537T7.75 1h8.5q.325 0 .537.213T17 1.75q0 .325-.213.537t-.537.213h-2.5L13 4q2.975.075 4.988.813T20 8v9.5q0 1.475-1.012 2.488T16.5 21l.5.5q.425.425.2.963t-.825.537q-.175 0-.338-.063t-.287-.187L14 21h-4l-1.75 1.75q-.125.125-.288.188T7.626 23q-.575 0-.813-.537T7 21.5l.5-.5q-1.475 0-2.488-1.012T4 17.5Zm8 .5q.625 0 1.063-.438T13.5 16.5q0-.625-.438-1.063T12 15q-.625 0-1.063.438T10.5 16.5q0 .625.438 1.063T12 18Zm-6-6h12V9H6v3Z"/>`),
		g.Group(children),
	)
}
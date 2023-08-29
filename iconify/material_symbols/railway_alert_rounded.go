package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailwayAlertRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17q.625 0 1.063-.438T11.5 15.5q0-.625-.438-1.063T10 14q-.625 0-1.063.438T8.5 15.5q0 .625.438 1.063T10 17Zm-8-.5V7q0-2.025 1.963-3.013T10 3h1.65q-.5 1.2-.613 2.463T11.3 8H4v3h9.1q1 1 2.25 1.5T18 13v3.5q0 1.475-1.012 2.488T14.5 20l1.625 1.075q.3.2.2.563T15.85 22H4.15q-.375 0-.475-.363t.2-.562L5.5 20q-1.475 0-2.488-1.012T2 16.5ZM18 11q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11Zm0-4q.2 0 .35-.15t.15-.35v-3q0-.2-.15-.35T18 3q-.2 0-.35.15t-.15.35v3q0 .2.15.35T18 7Zm0 2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T18 8q-.2 0-.35.15t-.15.35q0 .2.15.35T18 9Z"/>`),
		g.Group(children),
	)
}
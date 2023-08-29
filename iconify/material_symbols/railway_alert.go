package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailwayAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11Zm-.5-4h1V3h-1v4Zm.5 2q.2 0 .35-.15t.15-.35q0-.2-.15-.35T18 8q-.2 0-.35.15t-.15.35q0 .2.15.35T18 9Zm-8 8q.625 0 1.063-.438T11.5 15.5q0-.625-.438-1.063T10 14q-.625 0-1.063.438T8.5 15.5q0 .625.438 1.063T10 17Zm-6 5v-1l1.5-1q-1.475 0-2.488-1.012T2 16.5V7q0-2.025 1.963-3.013T10 3h1.65q-.3.725-.475 1.475T11 6q0 .525.075 1.012T11.3 8H4v3h9.1q1 1 2.25 1.5T18 13v3.5q0 1.475-1.012 2.488T14.5 20l1.5 1v1H4Z"/>`),
		g.Group(children),
	)
}
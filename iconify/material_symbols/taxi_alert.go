package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21v-8l2.1-6q.15-.45.537-.725T5.5 6H8V4h3.3q-.3.975-.3 2t.3 2H5.85L4.8 11h8.3q.975.975 2.25 1.488T18 13q.5 0 .963-.063t.937-.187l.1.25v8q0 .425-.287.713T19 22h-1q-.425 0-.713-.288T17 21v-1H5v1q0 .425-.288.713T4 22H3Zm3.5-5q.625 0 1.063-.438T8 15.5q0-.625-.438-1.063T6.5 14q-.625 0-1.063.438T5 15.5q0 .625.438 1.063T6.5 17ZM18 11q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11Zm-2.5 6q.625 0 1.063-.438T17 15.5q0-.625-.438-1.063T15.5 14q-.625 0-1.063.438T14 15.5q0 .625.438 1.063T15.5 17ZM18 9q.2 0 .35-.15t.15-.35q0-.2-.15-.35T18 8q-.2 0-.35.15t-.15.35q0 .2.15.35T18 9Zm-.5-2h1V3h-1v4Z"/>`),
		g.Group(children),
	)
}
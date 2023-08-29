package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DepartureBoardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 17q.625 0 1.063-.438T8 15.5q0-.625-.438-1.063T6.5 14q-.625 0-1.063.438T5 15.5q0 .625.438 1.063T6.5 17Zm7 0q.625 0 1.063-.438T15 15.5q0-.625-.438-1.063T13.5 14q-.625 0-1.063.438T12 15.5q0 .625.438 1.063T13.5 17ZM6 20v2H3v-3.05q-.45-.5-.725-1.113T2 16.5V7q0-.75.288-1.513t1.275-1.362q.987-.6 2.912-.912t5.175-.163q-.475 1.2-.6 2.45T11.3 8H4v3h9.1q.95.95 2.212 1.475T18 13v3.5q0 .725-.275 1.338T17 18.95V22h-3v-2H6Zm12-9q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11Zm1.65-2.6l.75-.75l-1.9-1.85V3h-1v3.2l2.15 2.2Z"/>`),
		g.Group(children),
	)
}
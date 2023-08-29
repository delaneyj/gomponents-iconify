package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoTransferOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19v.5q0 .625-.438 1.063T6.5 21q-.625 0-1.063-.438T5 19.5v-1.55q-.45-.5-.725-1.113T4 15.5V6.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3l-3.475-3.45h2.85v.3q0 .65-.462 1.113T17.575 21q-.65 0-1.113-.463T16 19.425V19H8Zm0-2h6.175l-5-5H6v3q0 .825.588 1.413T8 17Zm11.725-.15L18 15.125V12h-3.125l-2-2H18V7H9.875l-2-2h9.775q-.375-.425-1.613-.713T12.05 4q-1.775 0-2.888.15T7.4 4.525L5.875 3q.975-.5 2.488-.75T12 2q4.3 0 6.15.925T20 6v9.5q0 .35-.075.688t-.2.662ZM8.5 16q.625 0 1.063-.438T10 14.5q0-.625-.438-1.063T8.5 13q-.625 0-1.063.438T7 14.5q0 .625.438 1.063T8.5 16ZM6 10h1.175L6 8.825V10Zm1.875-5h9.775h-9.775Zm1.3 7Zm5.7 0Z"/>`),
		g.Group(children),
	)
}
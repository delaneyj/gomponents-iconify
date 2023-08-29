package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageInfoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.75 20.25q-1.575 0-2.663-1.088T14 16.5q0-1.575 1.088-2.663t2.662-1.087q1.575 0 2.663 1.088T21.5 16.5q0 1.575-1.088 2.663T17.75 20.25Zm0-2q.725 0 1.238-.513T19.5 16.5q0-.725-.513-1.238t-1.237-.512q-.725 0-1.238.513T16 16.5q0 .725.513 1.238t1.237.512ZM4 17.5v-2h8v2H4Zm2.25-6.25q-1.575 0-2.663-1.088T2.5 7.5q0-1.575 1.088-2.663T6.25 3.75q1.575 0 2.663 1.088T10 7.5q0 1.575-1.088 2.663T6.25 11.25Zm0-2q.725 0 1.238-.513T8 7.5q0-.725-.513-1.238T6.25 5.75q-.725 0-1.238.513T4.5 7.5q0 .725.513 1.238t1.237.512ZM12 8.5v-2h8v2h-8Zm5.75 8Zm-11.5-9Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFlashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-2.5-2.5q-.1.5-.513.85T16.5 22h-13q-.625 0-1.063-.438T2 20.5v-9.6q0-.625.438-1.063T3.5 9.4h2.65l.2-.225L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM4 20h12v-1.175l-2.575-2.575q-.275 1.2-1.213 1.975T10 19q-1.45 0-2.475-1.025T6.5 15.5q0-1.275.775-2.212t1.975-1.213L7.725 10.55L7 11.4H4V20Zm14-4.875l-2-2V11.4h-1.725l-1.7-1.7l-1.7-1.7h1.7l1.275 1.4h2.65q.625 0 1.063.437T18 10.9v4.225ZM19 11V7h-1V2h4l-1.6 3.6H22L19 11Zm-3 2.125Zm-4.125 1.55ZM10 17q.625 0 1.063-.438T11.5 15.5q0-.625-.438-1.063T10 14q-.625 0-1.063.438T8.5 15.5q0 .625.438 1.063T10 17Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DinnerDiningSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 21l-2-2h20l-2 2H4Zm-1-3q.15-.45.4-.85t.6-.75V9H3V7.5h1v-.75H3v-1.5h1V4.5H3V3h9v2.25h9v1.5h-9V9H8v6.1q.375.05.713.15t.637.3Q9.95 14 11.338 13t3.162-1q2.425 0 4.05 1.775T19.95 18H3ZM8 5.25h2V4.5H8v.75ZM8 7.5h2v-.75H8v.75ZM5.5 5.25h1V4.5h-1v.75Zm0 2.25h1v-.75h-1v.75Zm0 7.85q.225-.125.475-.187t.525-.113V9h-1v6.35Z"/>`),
		g.Group(children),
	)
}
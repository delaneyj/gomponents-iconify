package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taunt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.45 14.475q-.6.325-1.275.138T3.15 13.85q-.425-.725-.15-1.5t1.075-1.05L12.5 8.5l.9 1.775l-7.95 4.2ZM6 21v-5.675l8.725-4.6l-.425-.875L20 7l.9 1.8L14 14v7H6Zm1.5-11q-1.45 0-2.475-1.025T4 6.5q0-1.45 1.025-2.475T7.5 3q1.45 0 2.475 1.025T11 6.5q0 1.45-1.025 2.475T7.5 10Z"/>`),
		g.Group(children),
	)
}
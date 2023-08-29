package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18h1v-4q.625 0 1.063-.438T12.5 12.5v-3h-1v3H11v-3h-1v3h-.5v-3h-1v3q0 .625.438 1.063T10 14v4Zm4 0h1V9.5q-.825 0-1.413.588T13 11.5v3h1V18ZM4 21V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}
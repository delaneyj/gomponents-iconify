package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMealsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.575l-1.575-1.55h.8q0 .425-.287.713t-.713.287q-.425 0-.713-.287T17 21.024v-1.2L1.375 4.2q-.3-.3-.288-.713t.313-.712q.3-.3.713-.3t.712.3l18.375 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM19 16.125l-2.075-2.075L14 11.125V7q0-2.075 1.463-3.538T19 2v14.125Zm-7-7l-2-2V3q0-.425.288-.713T11 2q.425 0 .713.288T12 3v6.125Zm-3-3l-2-2V3q0-.425.288-.713T8 2q.425 0 .713.288T9 3v3.125Zm-3-3L4.875 2q.475 0 .8.325t.325.8ZM7 21v-8.15q-1.275-.35-2.138-1.4T4 9V3.975l2 2V9h1V6.975l2 2V9h.025l2.25 2.25q-.4.575-.988.988T9 12.85V21q0 .425-.288.713T8 22q-.425 0-.713-.288T7 21Z"/>`),
		g.Group(children),
	)
}
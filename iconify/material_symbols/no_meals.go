package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMeals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L19 21.825V22h-2v-2.175L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM19 16.125l-2.075-2.075L14 11.125V7q0-2.075 1.463-3.538T19 2v14.125Zm-7-7l-2-2V2h2v7.125Zm-3-3l-2-2V2h2v4.125Zm-3-3L4.875 2H6v1.125ZM7 22v-9.15q-1.275-.35-2.138-1.4T4 9V3.975l2 2V9h1V6.975l2.025 2l2.25 2.275q-.4.575-.988.988T9 12.85V22H7Z"/>`),
		g.Group(children),
	)
}
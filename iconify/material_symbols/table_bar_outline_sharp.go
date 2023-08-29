package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 20l2-5h3v-4.025Q7.175 10.85 4.587 9.85T2 7.5q0-1.45 2.925-2.475T12 4q4.175 0 7.088 1.025T22 7.5q0 1.35-2.588 2.35T13 10.975V15h3l2 5h-2l-1.2-3H9.2L8 20H6Zm6-11q2.425 0 4.575-.425t3.15-1.075q-1-.65-3.15-1.075T12 6q-2.425 0-4.575.425T4.275 7.5q1 .65 3.15 1.075T12 9Zm0-1.5Z"/>`),
		g.Group(children),
	)
}
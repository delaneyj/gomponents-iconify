package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.2 17l-.975 2.425q-.125.275-.35.425t-.5.15q-.5 0-.787-.413t-.088-.862l1-2.475q.225-.575.725-.913T9.35 15H11v-4.025Q7.175 10.85 4.587 9.85T2 7.5q0-1.45 2.925-2.475T12 4q4.175 0 7.088 1.025T22 7.5q0 1.35-2.588 2.35T13 10.975V15h1.65q.6 0 1.113.338t.737.912l1 2.475q.1.225.063.45t-.163.413q-.125.187-.325.3t-.45.112q-.275 0-.5-.15t-.35-.425L14.8 17H9.2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-1.65 0-2.825-1.175T2 18v-5.4q0-.775.363-1.437t1.012-1.088l8.25-5.35l.5.5q.5.5.588 1.188T12.425 7.7L11.3 9.5H19q.425 0 .713.288T20 10.5q0 .425-.288.713T19 11.5h-7V13h9q.425 0 .713.288T22 14q0 .425-.288.713T21 15h-9v1.5h8q.425 0 .713.288T21 17.5q0 .425-.288.713T20 18.5h-8V20h6q.425 0 .713.288T19 21q0 .425-.288.713T18 22H6Zm8.975-14q.125-1.025-.013-1.613t-.737-1.337q-.775-1-1.037-1.925T13.1 1H15q-.2 1.275.125 1.875t.9 1.4q.65.875.85 1.738T16.9 8h-1.925ZM19 8q.125-1.025-.025-1.613t-.75-1.337q-.775-1-1.025-1.925T17.125 1H19q-.2 1.275.125 1.875t.9 1.4q.65.875.85 1.738T20.9 8H19Z"/>`),
		g.Group(children),
	)
}
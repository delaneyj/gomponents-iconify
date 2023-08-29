package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 8q-.975 0-1.738-.763T15 5.5q0-.875.575-1.938T17.5 1q1.35 1.5 1.925 2.563T20 5.5q0 .975-.763 1.738T17.5 8ZM6 22q-1.65 0-2.825-1.175T2 18v-5.4q0-.775.363-1.437t1.012-1.088l8.25-5.35l.5.5q.5.5.588 1.188T12.425 7.7L11.3 9.5H19q.425 0 .713.288T20 10.5q0 .425-.288.713T19 11.5H7.7l1.925-3.1l-5.175 3.35q-.225.125-.338.35T4 12.6V18q0 .825.587 1.413T6 20h12q.425 0 .713.288T19 21q0 .425-.288.713T18 22H6Zm6-7v-2h9q.425 0 .713.288T22 14q0 .425-.288.713T21 15h-9Zm0 3.5v-2h8q.425 0 .713.288T21 17.5q0 .425-.288.713T20 18.5h-8ZM8 15Z"/>`),
		g.Group(children),
	)
}
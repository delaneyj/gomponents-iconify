package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-1.65 0-2.825-1.175T2 18v-5.4q0-.775.363-1.437t1.012-1.088l7.65-4.975q.275-.175.575-.137t.525.262q.5.5.588 1.188T12.425 7.7L11.3 9.5H19q.425 0 .713.288T20 10.5q0 .425-.288.713T19 11.5h-7V13h9q.425 0 .713.288T22 14q0 .425-.288.713T21 15h-9v1.5h8q.425 0 .713.288T21 17.5q0 .425-.288.713T20 18.5h-8V20h6q.425 0 .713.288T19 21q0 .425-.288.713T18 22H6ZM17.5 8q-.975 0-1.738-.763T15 5.5q0-1.125.725-2.188t1.4-1.862q.15-.2.375-.2t.375.2q.675.8 1.4 1.863T20 5.5q0 .975-.763 1.738T17.5 8Z"/>`),
		g.Group(children),
	)
}
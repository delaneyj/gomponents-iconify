package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LimboXEightySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 5.3h35.4v23.8H6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.7 8h30v18.4h-30Zm10 21.1l-4 4.9h18l-4-4.9M7.2 37l-2.3 4.5h37.6l-3-4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 37c-2.906-6.762 5.657-.81 5.787-4.043c.122-3.039-7.959-1.947-11.19-1.957"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.2 10H9.9q-1.575 0-2.738 1T6 13.5Q6 15 7.163 16T9.9 17H16q.425 0 .713.288T17 18q0 .425-.288.713T16 19H9.9q-2.425 0-4.163-1.575T4 13.5q0-2.35 1.738-3.925T9.9 8h6.3l-1.9-1.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l3.6 3.6q.15.15.213.325t.062.375q0 .2-.063.375T19.3 9.7l-3.6 3.6q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.9-1.9Z"/>`),
		g.Group(children),
	)
}
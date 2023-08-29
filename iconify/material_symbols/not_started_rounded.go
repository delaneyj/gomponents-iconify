package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotStartedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16q.425 0 .713-.288T10 15V9q0-.425-.288-.713T9 8q-.425 0-.713.288T8 9v6q0 .425.288.713T9 16Zm4.175-.775l3.9-2.6q.35-.225.35-.625t-.35-.625l-3.9-2.6q-.375-.25-.775-.038T12 9.4v5.2q0 .45.4.663t.775-.038ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
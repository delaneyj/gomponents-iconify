package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandGestureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5.025q0-1.275-.875-2.15T18.975 2V.5q1.875 0 3.2 1.325t1.325 3.2H22ZM6 23q-2.075 0-3.538-1.463T1 18h1.5q0 1.45 1.025 2.475T6 21.5V23Zm4.05 0q-.75 0-1.4-.338T7.575 21.7L1.2 12.375l.6-.575q.475-.475 1.125-.55t1.175.3L7 13.575V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v13.425l-3.7-2.6l3.925 5.725q.125.2.35.325t.475.125H17q.825 0 1.413-.588T19 19V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v14q0 1.65-1.175 2.825T17 23h-6.95ZM11 12V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v10h-2Zm4 0V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v9h-2Zm-2.85 4.5Z"/>`),
		g.Group(children),
	)
}
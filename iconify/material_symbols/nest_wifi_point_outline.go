package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWifiPointOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h14V9q0-1.65-1.175-2.825T15 5H9Q7.35 5 6.175 6.175T5 9v6Zm4 6q-2.5 0-4.25-1.75T3 15V9q0-2.5 1.75-4.25T9 3h6q2.5 0 4.25 1.75T21 9v6q0 2.5-1.75 4.25T15 21H9Zm3-6Zm-3 4q0-.425.288-.713T10 18q.425 0 .713.288T11 19h2q0-.425.288-.713T14 18q.425 0 .713.288T15 19q1.125 0 2.025-.55T18.45 17H17q0 .425-.288.713T16 18q-.425 0-.713-.288T15 17h-2q0 .425-.288.713T12 18q-.425 0-.713-.288T11 17H9q0 .425-.288.713T8 18q-.425 0-.713-.288T7 17H5.55q.525.9 1.425 1.45T9 19Z"/>`),
		g.Group(children),
	)
}
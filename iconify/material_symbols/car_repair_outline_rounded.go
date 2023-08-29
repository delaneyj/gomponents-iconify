package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRepairOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21v-2H5q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h14q.425 0 .713.288T20 18q0 .425-.288.713T19 19h-6v2q0 .425-.288.713T12 22ZM9 11.5q.425 0 .713-.288T10 10.5q0-.425-.288-.713T9 9.5q-.425 0-.713.288T8 10.5q0 .425.288.713T9 11.5Zm6 0q.425 0 .713-.288T16 10.5q0-.425-.288-.713T15 9.5q-.425 0-.713.288T14 10.5q0 .425.288.713T15 11.5ZM6 16q-.425 0-.713-.288T5 15V8.925q0-.175.025-.338t.075-.312L6.45 4.35q.2-.6.725-.975T8.35 3h7.3q.65 0 1.175.375t.725.975l1.35 3.925q.05.15.075.313t.025.337V15q0 .425-.287.713T18 16q-.425 0-.713-.288T17 15v-1H7v1q0 .425-.288.713T6 16Zm1.65-9h8.7l-.7-2h-7.3l-.7 2ZM7 9v3v-3Zm0 3h10V9H7v3Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRepairOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-3H4v-2h16v2h-7v3h-2ZM9 11.5q.425 0 .713-.288T10 10.5q0-.425-.288-.713T9 9.5q-.425 0-.713.288T8 10.5q0 .425.288.713T9 11.5Zm6 0q.425 0 .713-.288T16 10.5q0-.425-.288-.713T15 9.5q-.425 0-.713.288T14 10.5q0 .425.288.713T15 11.5ZM5 8.6l1.65-4.8q.125-.35.413-.575T7.7 3h8.6q.35 0 .638.225t.412.575L19 8.6v6.6q0 .35-.225.575T18.2 16h-.4q-.35 0-.575-.225T17 15.2V14H7v1.2q0 .35-.225.575T6.2 16h-.4q-.35 0-.575-.225T5 15.2V8.6ZM7.65 7h8.7l-.7-2h-7.3l-.7 2ZM7 9v3v-3Zm0 3h10V9H7v3Z"/>`),
		g.Group(children),
	)
}
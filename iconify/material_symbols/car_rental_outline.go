package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRentalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7q-1.25 0-2.125-.875T5 4q0-1.25.875-2.125T8 1q.95 0 1.725.563T10.85 3H19v2h-1v2h-2V5h-5.15q-.35.875-1.125 1.438T8 7Zm0-2q.425 0 .713-.288T9 4q0-.425-.288-.713T8 3q-.425 0-.713.288T7 4q0 .425.288.713T8 5Zm1 12.5q.425 0 .713-.288T10 16.5q0-.425-.288-.713T9 15.5q-.425 0-.713.288T8 16.5q0 .425.288.713T9 17.5Zm6 0q.425 0 .713-.288T16 16.5q0-.425-.288-.713T15 15.5q-.425 0-.713.288T14 16.5q0 .425.288.713T15 17.5ZM5 14.6l1.65-4.8q.125-.35.413-.575T7.7 9h8.6q.35 0 .638.225t.412.575L19 14.6v6.6q0 .35-.225.575T18.2 22h-.4q-.35 0-.575-.225T17 21.2V20H7v1.2q0 .35-.225.575T6.2 22h-.4q-.35 0-.575-.225T5 21.2v-6.6ZM7.65 13h8.7l-.7-2h-7.3l-.7 2ZM7 15v3v-3Zm0 3h10v-3H7v3Z"/>`),
		g.Group(children),
	)
}
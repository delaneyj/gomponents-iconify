package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.825 11l3.9 3.9q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288l-4.6-4.6q-.15-.15-.213-.325T2.426 11q0-.2.063-.375T2.7 10.3l4.6-4.6q.275-.275.688-.275T8.7 5.7q.3.3.3.713t-.3.712L4.825 11Zm6 1l2.9 2.9q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288l-4.6-4.6q-.15-.15-.213-.325T7.425 11q0-.2.063-.375T7.7 10.3l4.6-4.6q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712L10.825 10H17q2.075 0 3.538 1.463T22 15v3q0 .425-.288.713T21 19q-.425 0-.713-.288T20 18v-3q0-1.25-.875-2.125T17 12h-6.175Z"/>`),
		g.Group(children),
	)
}
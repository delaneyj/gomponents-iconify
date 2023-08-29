package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateVerticalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.3 19.3l-2.1-2.1q-.3-.3-.287-.713t.312-.712Q3.5 15.5 3.9 15.5t.7.275l.4.375V4q0-.425.288-.713T6 3q.425 0 .713.288T7 4v12.15l.4-.375q.3-.275.713-.263t.687.288q.275.275.275.7t-.275.7l-2.1 2.1q-.3.3-.7.3t-.7-.3Zm6.875-3.3q-.475 0-.738-.388t-.087-.812l3.325-8.9q.15-.425.513-.663T16 5q.45 0 .813.238t.512.662l3.325 8.9q.175.425-.088.813t-.737.387q-.275 0-.5-.163T19 15.4l-.8-2.2h-4.4l-.8 2.2q-.1.275-.325.438t-.5.162Zm2.175-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h2q.425 0 .713.288T10 18q0 .425-.288.713T9 19H6q-.425 0-.713-.288T5 18v-3q0-.425.288-.713T6 14q.425 0 .713.288T7 15v2ZM7 7v2q0 .425-.288.713T6 10q-.425 0-.713-.288T5 9V6q0-.425.288-.713T6 5h3q.425 0 .713.288T10 6q0 .425-.288.713T9 7H7Zm10 10v-2q0-.425.288-.713T18 14q.425 0 .713.288T19 15v3q0 .425-.288.713T18 19h-3q-.425 0-.713-.288T14 18q0-.425.288-.713T15 17h2Zm0-10h-2q-.425 0-.713-.288T14 6q0-.425.288-.713T15 5h3q.425 0 .713.288T19 6v3q0 .425-.288.713T18 10q-.425 0-.713-.288T17 9V7Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExitRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16H6q-.425 0-.713-.288T5 15q0-.425.288-.713T6 14h3q.425 0 .713.288T10 15v3q0 .425-.288.713T9 19q-.425 0-.713-.288T8 18v-2Zm8 0v2q0 .425-.288.713T15 19q-.425 0-.713-.288T14 18v-3q0-.425.288-.713T15 14h3q.425 0 .713.288T19 15q0 .425-.288.713T18 16h-2ZM8 8V6q0-.425.288-.713T9 5q.425 0 .713.288T10 6v3q0 .425-.288.713T9 10H6q-.425 0-.713-.288T5 9q0-.425.288-.713T6 8h2Zm8 0h2q.425 0 .713.288T19 9q0 .425-.288.713T18 10h-3q-.425 0-.713-.288T14 9V6q0-.425.288-.713T15 5q.425 0 .713.288T16 6v2Z"/>`),
		g.Group(children),
	)
}
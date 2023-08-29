package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14v6q0 .425-.288.713T6 21q-.425 0-.713-.288T5 20V5q0-.425.288-.713T6 4h7.175q.35 0 .625.225t.35.575L14.4 6H19q.425 0 .713.288T20 7v8q0 .425-.288.713T19 16h-5.175q-.35 0-.625-.225t-.35-.575L12.6 14H7Zm7.65 0H18V8h-4.425q-.35 0-.625-.225T12.6 7.2L12.35 6H7v6h6.425q.35 0 .625.225t.35.575l.25 1.2Zm-2.15-4Z"/>`),
		g.Group(children),
	)
}
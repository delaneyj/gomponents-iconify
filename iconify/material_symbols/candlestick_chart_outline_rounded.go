package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandlestickChartOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-.425 0-.713-.288T7 19v-1H6q-.425 0-.713-.288T5 17V7q0-.425.288-.713T6 6h1V5q0-.425.288-.713T8 4q.425 0 .713.288T9 5v1h1q.425 0 .713.288T11 7v10q0 .425-.288.713T10 18H9v1q0 .425-.288.713T8 20Zm-1-4h2V8H7v8Zm9 4q-.425 0-.713-.288T15 19v-4h-1q-.425 0-.713-.288T13 14V9q0-.425.288-.713T14 8h1V5q0-.425.288-.713T16 4q.425 0 .713.288T17 5v3h1q.425 0 .713.288T19 9v5q0 .425-.288.713T18 15h-1v4q0 .425-.288.713T16 20Zm-1-7h2v-3h-2v3Zm-7-1Zm8-.5Z"/>`),
		g.Group(children),
	)
}
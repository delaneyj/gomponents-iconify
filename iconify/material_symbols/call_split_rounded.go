package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallSplitRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7.4V9q0 .425-.288.713T5 10q-.425 0-.713-.288T4 9V5q0-.425.288-.713T5 4h4q.425 0 .713.288T10 5q0 .425-.288.713T9 6H7.4l5.025 5.025q.275.275.425.638t.15.762V19q0 .425-.287.713T12 20q-.425 0-.713-.288T11 19v-6.6l-5-5Zm12 0l-2.45 2.475q-.3.3-.713.3t-.712-.3q-.3-.3-.3-.725t.3-.725L16.6 6H15q-.425 0-.712-.288T14 5q0-.425.288-.713T15 4h4q.425 0 .713.288T20 5v4q0 .425-.288.713T19 10q-.425 0-.713-.288T18 9V7.4Z"/>`),
		g.Group(children),
	)
}
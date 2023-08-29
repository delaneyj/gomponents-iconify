package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonRemoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 11q-.425 0-.713-.288T16 10q0-.425.288-.713T17 9h4q.425 0 .713.288T22 10q0 .425-.288.713T21 11h-4Zm-8 1q-1.65 0-2.825-1.175T5 8q0-1.65 1.175-2.825T9 4q1.65 0 2.825 1.175T13 8q0 1.65-1.175 2.825T9 12Zm-7 8q-.425 0-.713-.288T1 19v-1.8q0-.85.438-1.563T2.6 14.55q1.55-.775 3.15-1.163T9 13q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T17 17.2V19q0 .425-.288.713T16 20H2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-1.5 0-2.388-1.163T3.1 17.25l1.8-6.8q-.825-.525-1.363-1.425T3 7q0-1.65 1.175-2.825T7 3h8q1.125 0 1.7.95t.075 1.95l-2 4q-.275.5-.738.8T13 11h-2.025l-.275 1h.3q.425 0 .713.288T12 13v2q0 .425-.288.713T11 16H9.65l-.75 2.8q-.275.975-1.075 1.588T6 21ZM19 5l-.625-1.35L22 2l.6 1.375L19 5Zm3 7l-3.625-1.625L19 9l3.6 1.65L22 12Zm-3-4.25v-1.5h4v1.5h-4Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoofingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.3 12q-.35 0-.475-.325t.15-.55L10.65 4.2q.275-.275.638-.388T12 3.7q.35 0 .713.113t.637.387L16 6.6V5q0-.425.288-.713T17 4h1q.425 0 .713.288T19 5v4.3l2.025 1.825q.275.225.15.55T20.7 12h-.925q-.375 0-.725-.138t-.625-.387L12 5.7l-6.425 5.775q-.275.25-.625.388T4.225 12H3.3ZM9 19v-4q0-.425.288-.713T10 14h4q.425 0 .713.288T15 15v4q0 .425-.288.713T14 20h-4q-.425 0-.713-.288T9 19Z"/>`),
		g.Group(children),
	)
}
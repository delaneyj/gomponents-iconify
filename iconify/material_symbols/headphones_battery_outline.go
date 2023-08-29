package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesBatteryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18H4q-.825 0-1.413-.588T2 16v-4q0-2.5 1.75-4.25T8 6q2.5 0 4.25 1.75T14 12v4q0 .825-.588 1.413T12 18h-2v-5h2.5v-1q0-1.875-1.313-3.188T8 7.5q-1.875 0-3.188 1.313T3.5 12v1H6v5ZM18 7V6h2v1h1q.425 0 .713.288T22 8v9q0 .425-.288.713T21 18h-4q-.425 0-.713-.288T16 17V8q0-.425.288-.713T17 7h1Zm0 9h2h-2Zm0 0h2V9h-2v7Z"/>`),
		g.Group(children),
	)
}
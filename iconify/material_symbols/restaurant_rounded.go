package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestaurantRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 22q-.425 0-.713-.288T17 21v-7h-1q-.825 0-1.413-.588T14 12V7q0-2.075 1.463-3.538T19 2v19q0 .425-.288.713T18 22ZM8 22q-.425 0-.713-.288T7 21v-8.15q-1.275-.35-2.138-1.4T4 9V3q0-.425.288-.713T5 2q.425 0 .713.288T6 3v6h1V3q0-.425.288-.713T8 2q.425 0 .713.288T9 3v6h1V3q0-.425.288-.713T11 2q.425 0 .713.288T12 3v6q0 1.4-.863 2.45T9 12.85V21q0 .425-.288.713T8 22Z"/>`),
		g.Group(children),
	)
}
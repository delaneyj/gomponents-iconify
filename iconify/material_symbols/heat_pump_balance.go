package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpBalance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18.6q1.25 0 2.125-.875T10 15.6v-6q0-.425.288-.712T11 8.6q.425 0 .713.288T12 9.6v6q0 1.25.875 2.125T15 18.6q1.25 0 2.125-.875T18 15.6v-7q0-.425.288-.713T19 7.6h1.175L19 8.75l1.4 1.425L24 6.6L20.4 3L19 4.425L20.175 5.6H19q-1.25 0-2.125.875T16 8.6v7q0 .425-.288.713T15 16.6q-.425 0-.713-.288T14 15.6v-6q0-1.25-.875-2.125T11 6.6q-1.25 0-2.125.875T8 9.6v6q0 .425-.288.713T7 16.6q-.425 0-.713-.288T6 15.6v-9H4v9q0 1.25.875 2.125T7 18.6Zm-4 3q-.825 0-1.413-.588T1 19.6v-8h22v8q0 .825-.588 1.413T21 21.6H3Z"/>`),
		g.Group(children),
	)
}
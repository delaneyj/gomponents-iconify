package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm4.95 1.95q1-1 1.525-2.275T19 12q0-1.4-.525-2.675T16.95 7.05l-1.4 1.4q.725.725 1.087 1.637T17 12q0 1-.363 1.913T15.55 15.55l1.4 1.4Zm-9.9 0l1.4-1.4q-.725-.725-1.088-1.638T7 12q0-1 .363-1.913T8.45 8.45l-1.4-1.4q-1 1-1.525 2.275T5 12q0 1.4.525 2.675T7.05 16.95ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}
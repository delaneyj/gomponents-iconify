package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurroundSoundOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm4.95 1.95q1-1 1.525-2.275T19 12q0-1.4-.525-2.675T16.95 7.05l-1.4 1.4q.725.725 1.087 1.637T17 12q0 1-.363 1.913T15.55 15.55l1.4 1.4Zm-9.9 0l1.4-1.4q-.725-.725-1.088-1.638T7 12q0-1 .363-1.913T8.45 8.45l-1.4-1.4q-1 1-1.525 2.275T5 12q0 1.4.525 2.675T7.05 16.95ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
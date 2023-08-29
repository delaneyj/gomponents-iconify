package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepOutRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.25 0-2.125-.875T9 19q0-1.25.875-2.125T12 16q1.25 0 2.125.875T15 19q0 1.25-.875 2.125T12 22ZM11 5.825L9.1 7.7q-.275.275-.688.287T7.7 7.7q-.275-.275-.275-.7t.275-.7l3.6-3.6q.3-.3.7-.3t.7.3l3.6 3.6q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288L13 5.825V13q0 .425-.288.713T12 14q-.425 0-.713-.288T11 13V5.825Z"/>`),
		g.Group(children),
	)
}
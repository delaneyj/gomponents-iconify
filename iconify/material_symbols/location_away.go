package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationAway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21v-1.9q0-.525.263-.988t.712-.737q1.15-.675 2.413-1.025T16 16q1.35 0 2.613.35t2.412 1.025q.45.275.713.738T22 19.1V21H10Zm-8 0V9l8-6l5.375 4.05q-1.875.225-3.125 1.638T11 12q0 .775.213 1.463t.612 1.287q-.5.2-.962.413t-.913.487q-.9.525-1.425 1.463T8 19.1V21H2Zm14-6q-1.25 0-2.125-.875T13 12q0-1.25.875-2.125T16 9q1.25 0 2.125.875T19 12q0 1.25-.875 2.125T16 15Z"/>`),
		g.Group(children),
	)
}
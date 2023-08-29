package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14v-2h8v2h-8Zm-4 7l-3.175-2.85q-1.8-1.625-3.088-2.9t-2.125-2.4q-.837-1.125-1.225-2.175T1 8.475q0-2.35 1.575-3.913T6.5 3q1.3 0 2.475.537T11 5.075q.85-1 2.025-1.538T15.5 3q2.125 0 3.563 1.288T20.85 7.3q-.45-.175-.9-.262t-.875-.088q-2.525 0-4.3 1.763T13 13q0 1.3.525 2.463T15 17.45q-.475.425-1.238 1.088T12.45 19.7L11 21Z"/>`),
		g.Group(children),
	)
}
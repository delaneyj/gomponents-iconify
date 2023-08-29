package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeLineOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11q-2.075 0-3.538-1.463T1 6q0-2.1 1.463-3.55T6 1q2.1 0 3.55 1.45T11 6q0 2.075-1.45 3.538T6 11Zm0-2q1.275 0 2.138-.875T9 6q0-1.275-.863-2.138T6 3q-1.25 0-2.125.863T3 6q0 1.25.875 2.125T6 9Zm8 14v-9h9v9h-9Zm2-2h5v-5h-5v5ZM6 6Zm11.725 1.7L7.7 17.7q.125.3.212.625T8 19q0 1.25-.863 2.125T5 22q-1.25 0-2.125-.875T2 19q0-1.275.875-2.138T5 16q.35 0 .675.088t.625.212l10-10.025q-.125-.3-.213-.613T16 5q0-1.275.875-2.138T19 2q1.275 0 2.138.863T22 5q0 1.25-.863 2.125T19 8q-.35 0-.663-.088t-.612-.212Zm.775 10.8Z"/>`),
		g.Group(children),
	)
}
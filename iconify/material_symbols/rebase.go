package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rebase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.75 23.25l-1.4-1.425L11.175 20h-3.35q-.325.875-1.088 1.438T5 22q-1.25 0-2.125-.875T2 19q0-.975.563-1.738T4 16.175v-8.35Q3.125 7.5 2.562 6.737T2 5q0-1.25.875-2.125T5 2q.975 0 1.738.563T7.824 4h3.35L9.35 2.175L10.75.75L15 5l-4.25 4.25l-1.4-1.425L11.175 6h-3.35q-.225.65-.7 1.125T6 7.825v8.35q.65.225 1.125.7t.7 1.125h3.35L9.35 16.175l1.4-1.425L15 19l-4.25 4.25ZM19 22q-1.25 0-2.125-.875T16 19q0-1 .563-1.763T18 16.175v-8.35q-.875-.3-1.438-1.063T16 5q0-1.25.875-2.125T19 2q1.25 0 2.125.875T22 5q0 1-.563 1.763T20 7.825v8.35q.875.325 1.438 1.088T22 19q0 1.25-.875 2.125T19 22Z"/>`),
		g.Group(children),
	)
}
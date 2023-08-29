package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RebaseEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.75 9.25l-1.4-1.425L11.175 6h-3.35q-.225.65-.7 1.125T6 7.825v8.35q.875.325 1.438 1.088T8 19q0 1.25-.875 2.125T5 22q-1.25 0-2.125-.875T2 19q0-.975.563-1.738T4 16.175v-8.35Q3.125 7.5 2.562 6.737T2 5q0-1.25.875-2.125T5 2q.975 0 1.738.563T7.824 4h3.35L9.35 2.175L10.75.75L15 5l-4.25 4.25ZM19 2q1.25 0 2.125.875T22 5q0 1.25-.875 2.125T19 8q-1.25 0-2.125-.875T16 5q0-1.25.875-2.125T19 2Zm-3.6 10.4l4.225 4.225L14.25 22H10v-4.25l5.4-5.35Zm5.3 3.175L16.425 11.3l.75-.725q.6-.6 1.438-.6t1.412.6l1.4 1.425q.575.575.575 1.4t-.575 1.4l-.725.775Z"/>`),
		g.Group(children),
	)
}
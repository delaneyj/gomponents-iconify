package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.887 2.27l-1.568.78l1.569.782l.78 1.569l.781-1.57l1.57-.78l-1.57-.781l-.78-1.569l-.782 1.569ZM7.96 1.515l1.097 2.204l2.204 1.097l-2.204 1.097L7.96 8.117L6.863 5.913L4.659 4.816l2.204-1.097L7.96 1.515Zm9.28 1.887l5.148 5.149L7.298 23.64L2.15 18.491L17.24 3.402Zm-2.005 4.833l2.32 2.32l2.005-2.004l-2.32-2.32l-2.005 2.004Zm.906 3.735l-2.32-2.32l-8.842 8.841l2.32 2.32l8.842-8.841Z"/>`),
		g.Group(children),
	)
}
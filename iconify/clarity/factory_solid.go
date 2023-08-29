package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactorySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityFactorySolid0" fill="currentColor" d="M32.45 8.44L22 15.3V9.51a1 1 0 0 0-1.63-.78L14.07 14H10V4.06L4 2.71V14H2v17a1 1 0 0 0 1 1h30a1 1 0 0 0 1-1V9.27a1 1 0 0 0-1.55-.83ZM14 29H6v-2h8Zm0-4H6v-2h8Zm0-4H6v-2h8Zm8 8h-2v-3h2Zm0-6h-2v-3h2Zm4 6h-2v-3h2Zm0-6h-2v-3h2Zm4 6h-2v-3h2Zm0-6h-2v-3h2Z"/>`),
		g.Group(children),
	)
}
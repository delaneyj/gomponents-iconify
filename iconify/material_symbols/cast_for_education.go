package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastForEducation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 10.6l4-2.3L16 6l-4 2.3l4 2.3Zm0 3.05l2.75-1.6v-2.1L16 11.55l-2.75-1.6v2.1l2.75 1.6ZM2 20v-3q1.25 0 2.125.875T5 20H2Zm5 0q0-2.075-1.463-3.538T2 15v-2q2.925 0 4.963 2.038T9 20H7Zm4 0q0-1.875-.713-3.513t-1.925-2.85q-1.212-1.212-2.85-1.924T2 11V9q2.275 0 4.275.863t3.5 2.362q1.5 1.5 2.363 3.5T13 20h-2Zm4 0q0-2.7-1.025-5.063t-2.788-4.124Q9.425 9.05 7.064 8.024T2 7V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20h-5Z"/>`),
		g.Group(children),
	)
}
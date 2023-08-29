package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RiceBowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22v-1.75Q5.375 19.2 3.687 17T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.8-1.688 5T16 20.25V22H8Zm2-10h4V4.25q-.5-.125-1-.188T12 4q-.5 0-1 .063t-1 .187V12Zm-6 0h4V5.075Q6.125 6.15 5.062 8T4 12Zm12 0h4q0-2.15-1.063-4T16 5.075V12Z"/>`),
		g.Group(children),
	)
}
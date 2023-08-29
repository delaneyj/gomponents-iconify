package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Productivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8Q8.35 8 7.175 6.825T6 4q0-1.65 1.175-2.825T10 0q1.65 0 2.825 1.175T14 4q0 1.65-1.175 2.825T10 8Zm6 0q-.275 0-.7-.063t-.7-.137q.675-.8 1.038-1.775T16 4q0-1.05-.363-2.025T14.6.2q.35-.125.7-.163T16 0q1.65 0 2.825 1.175T20 4q0 1.65-1.175 2.825T16 8Zm3 12q-2.075 0-3.538-1.463T14 15q0-2.1 1.463-3.55T19 10q2.1 0 3.55 1.45T24 15q0 2.075-1.45 3.538T19 20Zm-.7-2.75l3.525-3.55l-.7-.7l-2.825 2.825l-1.425-1.425l-.7.725L18.3 17.25ZM2 16v-2.8q0-.85.438-1.562T3.6 10.55q1.55-.775 3.15-1.163T10 9q1.125 0 2.225.175t2.2.55q-1.35 1.175-1.95 2.838t-.4 3.437H2Z"/>`),
		g.Group(children),
	)
}
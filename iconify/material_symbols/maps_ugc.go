package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapsUgc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 23l1.95-6.7q-.475-1.025-.713-2.1T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22q-1.125 0-2.2-.238t-2.1-.712L1 23Zm10-7h2v-3h3v-2h-3V8h-2v3H8v2h3v3Z"/>`),
		g.Group(children),
	)
}
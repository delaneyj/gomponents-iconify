package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapsUgcOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 23l1.95-6.7q-.475-1.025-.713-2.1T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22q-1.125 0-2.2-.238t-2.1-.712L1 23Zm2.95-2.95l3.2-.95q.35-.1.713-.075t.687.175q.8.4 1.675.6T12 20q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 .9.2 1.775t.6 1.675q.175.325.188.688t-.088.712l-.95 3.2ZM11 16h2v-3h3v-2h-3V8h-2v3H8v2h3v3Zm.975-3.975Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthAmerica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-2l-.7-.7q-.15-.15-.225-.325T11 18.6V13q-.825 0-1.413-.588T9 11v-1L5.875 6.875Q5 7.925 4.5 9.225T4 12q0 3.35 2.325 5.675T12 20Zm1-.05q2.975-.375 4.988-2.625T20 12q0-3.325-2.337-5.663T12 4q-1.1 0-2.113.288T8 5.074V7h3.55q.45 0 .863.2t.687.55l1.4 1.75H16q.425 0 .713.288T17 10.5v1.05q0 .225-.063.425t-.187.4L13 18v1.95Z"/>`),
		g.Group(children),
	)
}
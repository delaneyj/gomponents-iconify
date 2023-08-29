package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 17.5V16H13v-1.5h1.5V13H16v1.5h1.5V16H16v1.5h-1.5ZM7 12V6h1.5v2h2V6H12v6h-1.5V9.5h-2V12H7Zm6 0V6h3q.6 0 1.05.45t.45 1.05v3q0 .6-.45 1.05T16 12h-3Zm-2.5 3.5v-1h-2v1h2Zm-.05 3.5l-.85-2H8.5v2H7v-6h3.5q.625 0 1.063.438T12 14.5v1q0 .45-.263.813t-.637.587L12 19h-1.55Zm4.05-8.5H16v-3h-1.5v3ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-2q3.325 0 5.663-2.337T20 12q0-3.325-2.337-5.663T12 4Q8.675 4 6.337 6.337T4 12q0 3.325 2.337 5.663T12 20Zm0-8Z"/>`),
		g.Group(children),
	)
}
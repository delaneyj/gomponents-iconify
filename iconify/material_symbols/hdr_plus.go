package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 17.5H16V16h1.5v-1.5H16V13h-1.5v1.5H13V16h1.5v1.5ZM7 12h1.5V9.5h2V12H12V6h-1.5v2h-2V6H7v6Zm6 0h3q.6 0 1.05-.45t.45-1.05v-3q0-.6-.45-1.05T16 6h-3v6Zm-2.5 3.5h-2v-1h2v1Zm-.05 3.5H12l-.9-2.1q.375-.225.638-.588T12 15.5v-1q0-.625-.438-1.063T10.5 13H7v6h1.5v-2h1.1l.85 2Zm4.05-8.5v-3H16v3h-1.5ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
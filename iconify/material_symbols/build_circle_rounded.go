package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildCircleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.5 16.9l1.4-1.4q.15-.15.15-.35t-.15-.35l-3.4-3.425q.1-.275.15-.55t.05-.625q0-1.425-1.012-2.438T10.25 6.75q-.275 0-.525.025t-.5.125Q9 7 8.95 7.25t.125.425l1.85 1.85l-1.4 1.4l-1.85-1.85Q7.5 8.9 7.25 8.95t-.35.275q-.075.25-.113.5t-.037.525q0 1.425 1.012 2.438T10.2 13.7q.325 0 .613-.05t.562-.15l3.425 3.4q.15.15.35.15t.35-.15ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
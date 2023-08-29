package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorMedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.225 12.425l1.45 1.975q.15.2.4.2t.4-.2l1.45-1.975L14.35 14.4q.15.2.413.2t.412-.2l2.25-3.075q.25-.35.188-.75t-.413-.65q-.35-.25-.75-.188t-.65.413l-1.05 1.425L13.325 9.6q-.15-.2-.413-.2t-.412.2l-1.425 1.975L9.625 9.6q-.15-.2-.4-.2t-.4.2l-2.25 3.075q-.25.35-.188.75t.413.65q.35.25.75.188t.65-.413l1.025-1.425ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}
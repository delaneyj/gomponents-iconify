package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposurePlusOneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h2v-2q0-.425.288-.713T7 9q.425 0 .713.288T8 10v2h2q.425 0 .713.288T11 13q0 .425-.288.713T10 14H8v2q0 .425-.288.713T7 17q-.425 0-.713-.288T6 16v-2Zm9.75-5.95l-1.425 1.025q-.35.275-.788.188T12.85 8.8q-.225-.35-.15-.763t.425-.662L16.2 5.15q.1-.075.213-.112T16.65 5h.55q.35 0 .575.225T18 5.8v12.075q0 .475-.325.8t-.8.325q-.475 0-.8-.325t-.325-.8V8.05Z"/>`),
		g.Group(children),
	)
}
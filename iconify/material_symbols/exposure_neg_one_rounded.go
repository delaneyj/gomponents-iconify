package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposureNegOneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h5q.425 0 .713.288T10 13q0 .425-.288.713T9 14Zm6.75-5.95l-1.425 1.025q-.35.275-.788.188T12.85 8.8q-.225-.35-.15-.763t.425-.662L16.2 5.15q.1-.075.213-.112T16.65 5h.55q.35 0 .575.225T18 5.8v12.075q0 .475-.325.8t-.8.325q-.475 0-.8-.325t-.325-.8V8.05Z"/>`),
		g.Group(children),
	)
}
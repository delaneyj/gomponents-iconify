package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21v-3H6.625q-.3 0-.588-.125t-.487-.325l-1.5-1.5q-.225-.225-.325-.5t-.1-.55q0-.275.1-.55t.325-.5l1.5-1.5q.2-.2.488-.325T6.625 12H11v-2H5.5q-.625 0-1.063-.438T4 8.5v-3q0-.625.438-1.063T5.5 4H11V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v1h4.375q.3 0 .588.125t.487.325l1.5 1.5q.225.225.325.5t.1.55q0 .275-.1.55t-.325.5l-1.5 1.5q-.2.2-.488.325t-.587.125H13v2h5.5q.625 0 1.063.438T20 13.5v3q0 .625-.438 1.063T18.5 18H13v3q0 .425-.288.713T12 22Z"/>`),
		g.Group(children),
	)
}
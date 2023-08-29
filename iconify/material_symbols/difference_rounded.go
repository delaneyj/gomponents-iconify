package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DifferenceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 9v1q0 .425.288.713T13.5 11q.425 0 .713-.288T14.5 10V9h1q.425 0 .713-.288T16.5 8q0-.425-.288-.713T15.5 7h-1V6q0-.425-.288-.713T13.5 5q-.425 0-.713.288T12.5 6v1h-1q-.425 0-.713.288T10.5 8q0 .425.288.713T11.5 9h1Zm-1 6h4q.425 0 .713-.288T16.5 14q0-.425-.288-.713T15.5 13h-4q-.425 0-.713.288T10.5 14q0 .425.288.713T11.5 15ZM8 19q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h6.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V17q0 .825-.588 1.413T19 19H8Zm-4 4q-.825 0-1.413-.588T2 21V8q0-.425.288-.713T3 7q.425 0 .713.288T4 8v13h10q.425 0 .713.288T15 22q0 .425-.288.713T14 23H4Z"/>`),
		g.Group(children),
	)
}
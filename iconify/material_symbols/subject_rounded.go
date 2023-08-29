package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubjectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h8q.425 0 .713.288T14 18q0 .425-.288.713T13 19H5Zm0-4q-.425 0-.713-.288T4 14q0-.425.288-.713T5 13h14q.425 0 .713.288T20 14q0 .425-.288.713T19 15H5Zm0-4q-.425 0-.713-.288T4 10q0-.425.288-.713T5 9h14q.425 0 .713.288T20 10q0 .425-.288.713T19 11H5Zm0-4q-.425 0-.713-.288T4 6q0-.425.288-.713T5 5h14q.425 0 .713.288T20 6q0 .425-.288.713T19 7H5Z"/>`),
		g.Group(children),
	)
}
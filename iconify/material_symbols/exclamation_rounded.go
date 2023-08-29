package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q-.425 0-.713-.288T11 13V6q0-.425.288-.713T12 5q.425 0 .713.288T13 6v7q0 .425-.288.713T12 14Zm0 5q-.425 0-.713-.288T11 18q0-.425.288-.713T12 17q.425 0 .713.288T13 18q0 .425-.288.713T12 19Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapTextRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 12.5q-.425 0-.713-.288T4 11.5q0-.425.288-.713T5 10.5h12.25q1.575 0 2.663 1.088T21 14.25q0 1.575-1.088 2.663T17.25 18h-2.4l.55.55q.3.3.288.7t-.288.7q-.3.3-.713.313t-.712-.288L11.7 17.7q-.15-.15-.212-.325T11.425 17q0-.2.063-.375t.212-.325l2.275-2.275q.3-.3.713-.3t.712.3q.275.3.288.713t-.288.712l-.55.55h2.4q.725 0 1.238-.513T19 14.25q0-.725-.513-1.238T17.25 12.5H5ZM5 18q-.425 0-.713-.288T4 17q0-.425.288-.713T5 16h3q.425 0 .713.288T9 17q0 .425-.288.713T8 18H5ZM5 7q-.425 0-.713-.288T4 6q0-.425.288-.713T5 5h14q.425 0 .713.288T20 6q0 .425-.288.713T19 7H5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeArrowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 22q-2.7 0-4.6-1.9T7 15.5v-.575q-2.15-.35-3.575-2.013T2 9V4q0-.425.288-.713T3 3h2q0-.425.288-.713T6 2q.425 0 .713.288T7 3v2q0 .425-.288.713T6 6q-.425 0-.713-.288T5 5H4v4q0 1.65 1.175 2.825T8 13q1.65 0 2.825-1.175T12 9V5h-1q0 .425-.288.713T10 6q-.425 0-.713-.288T9 5V3q0-.425.288-.713T10 2q.425 0 .713.288T11 3h2q.425 0 .713.288T14 4v5q0 2.25-1.425 3.913T9 14.925v.575q0 1.875 1.313 3.188T13.5 20v2Zm1.5-4q-.425 0-.713-.288T14 17q0-.425.288-.713T15 16h3.175l-.875-.875q-.275-.3-.288-.713t.288-.712q.3-.3.7-.3t.7.3l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.3.3-.7.3t-.7-.3q-.3-.3-.288-.713t.288-.712l.875-.875H15Z"/>`),
		g.Group(children),
	)
}
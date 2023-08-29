package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryAlertRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v16q0 .425-.288.713T16 22H8Zm4-8q.425 0 .713-.288T13 13V9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9v4q0 .425.288.713T12 14Zm0 4q.425 0 .713-.288T13 17q0-.425-.288-.713T12 16q-.425 0-.713.288T11 17q0 .425.288.713T12 18Z"/>`),
		g.Group(children),
	)
}
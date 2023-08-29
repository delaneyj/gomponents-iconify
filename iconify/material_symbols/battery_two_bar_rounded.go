package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryTwoBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v16q0 .425-.288.713T16 22H8Zm1-6h6V6H9v10Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingTwentyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 14.1v2.75q0 .2.2.25t.275-.15L13.7 12.8q.1-.25-.038-.475t-.412-.225h-.75V9.35q0-.2-.2-.25t-.275.15L10.3 13.4q-.1.25.038.475t.412.225h.75ZM8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v16q0 .425-.288.713T16 22H8Zm1-2h6V6H9v14Zm0 0h6h-6Z"/>`),
		g.Group(children),
	)
}
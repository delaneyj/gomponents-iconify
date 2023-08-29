package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingFullOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 19h-1.55q-.3 0-.438-.263t.038-.512l2.5-3.575q.125-.15.288-.1t.162.25V17h1.55q.3 0 .438.263t-.038.512l-2.5 3.575q-.125.15-.288.1t-.162-.25V19ZM9 20Zm-1 2q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v7q-.525 0-1.025.088T15 12.35V6H9v14h2.35q.2.575.488 1.075t.687.925H8Z"/>`),
		g.Group(children),
	)
}
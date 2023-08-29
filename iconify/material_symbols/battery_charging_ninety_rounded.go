package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingNinetyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v7q-2.5.025-4.25 1.763T11 18q0 1.15.4 2.175T12.525 22H8ZM9 8h6V6H9v2Zm7.5 11h-1.55q-.3 0-.438-.263t.038-.512l2.5-3.575q.125-.15.288-.1t.162.25V17h1.55q.3 0 .438.263t-.038.512l-2.5 3.575q-.125.15-.288.1t-.162-.25V19Z"/>`),
		g.Group(children),
	)
}
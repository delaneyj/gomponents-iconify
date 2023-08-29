package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBottleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22L7 11l.475-.25q.2-.125.35-.3t.15-.425q0-.225-.1-.425t-.3-.3L7 9l1-4h3.5V4H10V2h5v2h-1.5v1H17l.975 3.95l-.575.3q-.2.1-.3.3t-.1.425q0 .25.15.438t.35.312l.475.225L17 22H8Z"/>`),
		g.Group(children),
	)
}
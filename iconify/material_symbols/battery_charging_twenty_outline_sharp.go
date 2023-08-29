package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingTwentyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22V4h3V2h4v2h3v18H7Zm2-2h6V6H9v14Zm0 0h6h-6Zm2.5-2l2.5-6h-1.5V8L10 14h1.5v4Z"/>`),
		g.Group(children),
	)
}
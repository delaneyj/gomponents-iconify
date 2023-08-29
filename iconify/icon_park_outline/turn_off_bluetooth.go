package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnOffBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M20.672 11.778V4l14.584 11.111l-7.179 4.445M32.5 35.974L20.675 44V29.628L32.5 35.974Z"/><path d="m4 12.5l40 23"/><path stroke-linejoin="round" d="m7.444 34l13.231-4.373"/></g>`),
		g.Group(children),
	)
}
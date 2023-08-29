package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineeringVehicle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M32 6H38"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 36H6V28H32V36H18"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 36V12H38.5L44 24V36H41"/><path fill="#2F88FF" stroke-linejoin="round" d="M4 8L26 8L26 22L7 22L4 8Z"/><circle cx="37" cy="38" r="4" fill="#2F88FF"/><circle cx="14" cy="38" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
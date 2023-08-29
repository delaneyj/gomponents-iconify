package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarberClippers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 8L38 8V17L33 24V36C33 36 33 44 24 44C15 44 15 36 15 36L15 24L10 17V8Z"/><path d="M15 4V10"/><path d="M21 4V10"/><path d="M27 4V10"/><rect width="6" height="10" x="21" y="28" fill="#2F88FF" rx="3"/><path d="M10 17H38"/><path d="M33 4V10"/></g>`),
		g.Group(children),
	)
}
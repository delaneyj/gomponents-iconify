package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WholeSiteAccelerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="22" cy="40" r="4" fill="#2F88FF"/><circle cx="26" cy="8" r="4" fill="#2F88FF"/><circle cx="36" cy="24" r="4" fill="#2F88FF"/><circle cx="12" cy="24" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 24L16 24"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 11L15 21"/><path d="M32.9998 27L24.9987 37"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 44L22.6875 15.5"/><path d="M36 44L25.3125 15.5"/><circle cx="24" cy="12" r="4" fill="#2F88FF"/><path d="M37.57 33C33.6618 35.5307 29.0024 37 23.9998 37C18.9973 37 14.3379 35.5307 10.4297 33"/><path d="M24 8V4"/></g>`),
		g.Group(children),
	)
}
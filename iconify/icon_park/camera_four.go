package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path d="M8 10V24C8 32.8366 15.1634 40 24 40V40C32.8366 40 40 32.8366 40 24V10"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 10H44"/><path fill="#2F88FF" stroke-linejoin="round" d="M24 30C27.3137 30 30 27.3137 30 24C30 20.6863 27.3137 18 24 18C20.6863 18 18 20.6863 18 24C18 27.3137 20.6863 30 24 30Z"/></g>`),
		g.Group(children),
	)
}
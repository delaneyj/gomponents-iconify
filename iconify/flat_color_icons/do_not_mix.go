package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotMix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M26.9 42H17v-9.9z"/><path fill="#00BCD4" d="M30 6v20.2L19.8 36.4l2.8 2.8L34 27.8V6z"/><path fill="#2196F3" d="M15.9 31H6v-9.9z"/><path fill="#2196F3" d="M20.2 14L8.8 25.4l2.8 2.8L21.8 18H41v-4z"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		g.Group(children),
	)
}